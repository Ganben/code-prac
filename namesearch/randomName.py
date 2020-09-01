# -*- coding: UTF-8 -*-

import random
import os
import ast

result = []
chardict = []
# chardict = []
#init standard char table
line1 = []
line2 = []
with open('normalchar2013.txt', 'r',encoding='utf-8') as fo:
    line1 = fo.read()
    line2 = fo.read()
for e in line1:
    chardict.append(e)
for e in line2:
    chardict.append(e)

print('loaded %s chars' % len(chardict))

def generate():
    with open("savedchars.txt", "r", encoding="utf-8") as fr:
        l = fr.readlines()
    
    n = len(l)
    for i in range(0, 10):
        k = random.randint(0, n-1)
        print("-%s ", l[k])
    




# for i in range(2):
#     for j in range(3):
#         key1 = random.randint(0, 6500-1)
#         key2 = random.randint(0, 3500-1)
#         result.append((chardict[key1], chardict[key2]))

# for k in result:
#     print(k[0], k[1])

def GetInput(user_message, var_type = str):
    while 1:
        # ask the user for an input
        str_input = input(user_message + ": ")
        # you dont need to cast a string!
        if len(str_input) > 0:
            if str_input == "exit":
                return "exit"
            return True
        try:

            input_type = type(ast.literal_eval(str_input))
        except:

            print("kept\n")
            return False
        if var_type == input_type or input_type == int:
            # return ast.literal_eval(str_input)
            print("deleted\n")
            return True

        else:
            print("kept(%s)\n" % input_type)
            return False

def saveStopChar(l):
    with open('stopchar.txt', 'a+') as fw:
        fw.writelines(["%s\n" % e for e in l])
    return

def saveProceededChartable(l):
    with open('savedchars.txt', 'a+', encoding="utf-8") as fw:
        fw.writelines(["%s\n" % e for e in l])
    return

def main():
    my_bool = False
    keepChars = []
    # my_str = ""
    # my_num = 0
    # my_bool = GetInput("Give me a Boolean", type(my_bool))
    # my_str = GetInput("Give me a String", type(my_str))
    # my_num = GetInput("Give me a Integer", type(my_num))
    with open('stopchar.txt', 'r') as fr:
        stopedNs = fr.readlines()
    
    n1 = 0
    if len(stopedNs) > 0:
        n1 = int(stopedNs[len(stopedNs)-1])+1
        print('load from %s' % n1)
    
    deletedNums = []
    for i in range(n1, n1+30):
        print("?(%d): %s" % (i, chardict[i]))
        r = GetInput("skip or any input for stop")
        if r == 'exit':
            break
        if r:

            deletedNums.append(i)
            print("deleted cache %d" % len(deletedNums))
        else:
            keepChars.append(chardict[i])
    saveStopChar(deletedNums)
    saveProceededChartable(keepChars)

    # if my_bool:
    #     print('"{}"'.format(my_str))
    #     print(my_str)
    # else:
    #     print(my_num * 2)

if __name__ == "__main__":
    # main()
    generate()