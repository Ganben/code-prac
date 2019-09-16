from __future__ import print_function
import paddle
import paddle.fluid as fluid
import numpy
import math
import sys

BATCH_SIZE = 20

train_reader = paddle.batch(
    paddle.reader.shuffle(
        paddle.dataset.uci_housing.train(), buf_size=500),
        batch_size=BATCH_SIZE)

test_reader = paddle.batch(
    paddle.reader.shuffle(
        paddle.dataset.uci_housing.test(), buf_size=500),
        batch_size=BATCH_SIZE)

x = fluid.layers.data(name='x', shape=[13], dtype='float32') # define shape and data type of input
y = fluid.layers.data(name='y', shape=[1], dtype='float32') # define shape and data type of output
y_predict = fluid.layers.fc(input=x, size=1, act=None) # fully connected layer connecting input and output

main_program = fluid.default_main_program() # get default/global main function
startup_program = fluid.default_startup_program() # get default/global launch program

cost = fluid.layers.square_error_cost(input=y_predict, label=y) # use label and output predicted data to estimate square error
avg_loss = fluid.layers.mean(cost) # compute mean value for square error and get mean los

#Clone main_program to get test_program
# operations of some operators are different between train and test. For example, batch_norm use parameter for_test to determine whether the program is for training or for testing.
#The api will not delete any operator, please apply it before backward and optimization.
test_program = main_program.clone(for_test=True)

sgd_optimizer = fluid.optimizer.SGD(learning_rate=0.001)
sgd_optimizer.minimize(avg_loss)

use_cuda = False
place = fluid.CUDAPlace(0) if use_cuda else fluid.CPUPlace() # define the execution space of executor

###executor can accept input program and add data input operator and result fetch operator based on feed map and fetch list. Use close() to close executor and call run(...) to run the program.
exe = fluid.Executor(place)

num_epochs = 100

def train_test(executor, program, reader, feeder, fetch_list):
    accumulated = 1 * [0]
    count = 0
    for data_test in reader():
        outs = executor.run(program=program,
                            feed=feeder.feed(data_test),
                            fetch_list=fetch_list)
        accumulated = [x_c[0] + x_c[1][0] for x_c in zip(accumulated, outs)] # accumulate loss value in the process of test
        count += 1 # accumulate samples in test dataset
    return [x_d / count for x_d in accumulated] # compute mean loss


