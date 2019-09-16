from __future__ import print_function
import paddle as paddle
import paddle.fluid as fluid
import six
import numpy
import math



EMBED_SIZE = 32 # embedding dimensions
HIDDEN_SIZE = 256 # hidden layer size
N = 5 # ngram size, here fixed 5
BATCH_SIZE = 100 # batch size
PASS_NUM = 100 # Training rounds

use_cuda = False # Set to True if trained with GPU

word_dict = paddle.dataset.imikolov.build_dict()
dict_size = len(word_dict)

def inference_program(words, is_sparse):

    embed_first = fluid.layers.embedding(
        input=words[0],
        size=[dict_size, EMBED_SIZE],
        dtype='float32',
        is_sparse=is_sparse,
        param_attr='shared_w')
    embed_second = fluid.layers.embedding(
        input=words[1],
        size=[dict_size, EMBED_SIZE],
        dtype='float32',
        is_sparse=is_sparse,
        param_attr='shared_w')
    embed_third = fluid.layers.embedding(
        input=words[2],
        size=[dict_size, EMBED_SIZE],
        dtype='float32',
        is_sparse=is_sparse,
        param_attr='shared_w')
    embed_fourth = fluid.layers.embedding(
        input=words[3],
        size=[dict_size, EMBED_SIZE],
        dtype='float32',
        is_sparse=is_sparse,
        param_attr='shared_w')

    concat_embed = fluid.layers.concat(
        input=[embed_first, embed_second, embed_third, embed_fourth], axis=1)
    hidden1 = fluid.layers.fc(input=concat_embed,
                              size=HIDDEN_SIZE,
                              act='sigmoid')
    predict_word = fluid.layers.fc(input=hidden1, size=dict_size, act='softmax')
    return predict_word

def train_program(predict_word):
    # The definition of'next_word' must be after the declaration of inference_program.
    # Otherwise the sequence of the train program input data becomes [next_word, firstw, secondw,
    #thirdw, fourthw], This is not true.
    next_word = fluid.layers.data(name='nextw', shape=[1], dtype='int64')
    cost = fluid.layers.cross_entropy(input=predict_word, label=next_word)
    avg_cost = fluid.layers.mean(cost)
    return avg_cost

def optimizer_func():
    return fluid.optimizer.AdagradOptimizer(
        learning_rate=3e-3,
        regularization=fluid.regularizer.L2DecayRegularizer(8e-4))

def train(if_use_cuda, params_dirname, is_sparse=True):
    place = fluid.CUDAPlace(0) if if_use_cuda else fluid.CPUPlace()

    train_reader = paddle.batch(
        paddle.dataset.imikolov.train(word_dict, N), BATCH_SIZE)
    test_reader = paddle.batch(
        paddle.dataset.imikolov.test(word_dict, N), BATCH_SIZE)

    first_word = fluid.layers.data(name='firstw', shape=[1], dtype='int64')
    second_word = fluid.layers.data(name='secondw', shape=[1], dtype='int64')
    third_word = fluid.layers.data(name='thirdw', shape=[1], dtype='int64')
    forth_word = fluid.layers.data(name='fourthw', shape=[1], dtype='int64')
    next_word = fluid.layers.data(name='nextw', shape=[1], dtype='int64')

    word_list = [first_word, second_word, third_word, forth_word, next_word]
    feed_order = ['firstw', 'secondw', 'thirdw', 'fourthw', 'nextw']

    main_program = fluid.default_main_program()
    star_program = fluid.default_startup_program()

    predict_word = inference_program(word_list, is_sparse)
    avg_cost = train_program(predict_word)
    test_program = main_program.clone(for_test=True)

    sgd_optimizer = optimizer_func()
    sgd_optimizer.minimize(avg_cost)

    exe = fluid.Executor(place)

    def train_test(program, reader):
        count = 0
        feed_var_list = [
            program.global_block().var(var_name) for var_name in feed_order
        ]
        feeder_test = fluid.DataFeeder(feed_list=feed_var_list, place=place)
        test_exe = fluid.Executor(place)
        accumulated = len([avg_cost]) * [0]
        for test_data in reader():
            avg_cost_np = test_exe.run(
                program=program,
                feed=feeder_test.feed(test_data),
                fetch_list=[avg_cost])
            accumulated = [
                x[0] + x[1][0] for x in zip(accumulated, avg_cost_np)
            ]
        count += 1
        return [x / count for x in accumulated]

    def train_loop():
        step = 0
        feed_var_list_loop = [
            main_program.global_block().var(var_name) for var_name in feed_order
        ]
        feeder = fluid.DataFeeder(feed_list=feed_var_list_loop, place=place)
        exe.run(star_program)
        for pass_id in range(PASS_NUM):
            for data in train_reader():
                avg_cost_np = exe.run(
                    main_program, feed=feeder.feed(data), fetch_list=[avg_cost])

                if step % 10 == 0:
                    outs = train_test(test_program, test_reader)

                    print("Step %d: Average Cost %f" % (step, outs[0]))

                    # The entire training process takes several hours if the average loss is less than 5.8,
                    # We think that the model has achieved good results and can stop training.
                    # Note 5.8 is a relatively high value, in order to get a better model, you can
                    # set the threshold here to be 3.5, but the training time will be longer.
                    if outs[0] < 5.8:
                        if params_dirname is not None:
                            fluid.io.save_inference_model(params_dirname, [
                                 'firstw', 'secondw', 'thirdw', 'fourthw'
                            ], [predict_word], exe)
                        return
                step += 1
                if math.isnan(float(avg_cost_np[0])):
                    sys.exit("got NaN loss, training failed.")

        raise AssertionError("Cost is too large {0:2.2}".format(avg_cost_np[0]))

    train_loop()

def infer(use_cuda, params_dirname=None):
    place = fluid.CUDAPlace(0) if use_cuda else fluid.CPUPlace()

    exe = fluid.Executor(place)

    inference_scope = fluid.core.Scope()
    with fluid.scope_guard(inference_scope):
        #Get the inference program using fluid.io.load_inference_model,
        #feed variable name by feed_target_names and fetch fetch_targets from scope
        [inferencer, feed_target_names,
        fetch_targets] = fluid.io.load_inference_model(params_dirname, exe)

        # Set the input and use 4 LoDTensor to represent 4 words. Each word here is an id,
        # Used to query the embedding table to get the corresponding word vector, so its shape size is [1].
        # recursive_sequence_lengths sets the length based on LoD, so it should all be set to [[1]]
        # Note that recursive_sequence_lengths is a list of lists
        data1 = numpy.asarray([[211]], dtype=numpy.int64)  # 'among'
        data2 = numpy.asarray([[6]], dtype=numpy.int64)  # 'a'
        data3 = numpy.asarray([[96]], dtype=numpy.int64)  # 'group'
        data4 = numpy.asarray([[4]], dtype=numpy.int64)  # 'of'
        lod = numpy.asarray([[1]], dtype=numpy.int64)

        first_word = fluid.create_lod_tensor(data1, lod, place)
        second_word = fluid.create_lod_tensor(data2, lod, place)
        third_word = fluid.create_lod_tensor(data3, lod, place)
        fourth_word = fluid.create_lod_tensor(data4, lod, place)

        assert feed_target_names[0] == 'firstw'
        assert feed_target_names[1] == 'secondw'
        assert feed_target_names[2] == 'thirdw'
        assert feed_target_names[3] == 'fourthw'

        # Construct the feed dictionary {feed_target_name: feed_target_data}
        # Prediction results are included in results
        results = exe.run(
            inferencer,
            feed={
                feed_target_names[0]: first_word,
                feed_target_names[1]: second_word,
                feed_target_names[2]: third_word,
                feed_target_names[3]: fourth_word
            },
            fetch_list=fetch_targets,
            return_numpy=False)

        print(numpy.array(results[0]))
        most_possible_word_index = numpy.argmax(results[0])
        print(most_possible_word_index)
        print([
            key for key, value in six.iteritems(word_dict)
            if value == most_possible_word_index
        ][0])

def main(use_cuda, is_sparse):
    if use_cuda and not fluid.core.is_compiled_with_cuda():
        return

    params_dirname = "word2vec.inference.model"

    train(
        if_use_cuda=use_cuda,
        params_dirname=params_dirname,
        is_sparse=is_sparse)

    infer(use_cuda=use_cuda, params_dirname=params_dirname)


main(use_cuda=use_cuda, is_sparse=True)
