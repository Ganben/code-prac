# 
from ward import test
import time
import datetime

def window(t):
    # return True if time in between 0000 2355,
    # return false if time between 2355 0000,
    # listening to manually triggered version
    # pf = rd.get('flag-clearing')
    # try:
    #     if pf is None:
    #         pf = -1
    #     else:
    #         pf = int(pf)
    # except:
    #     pf = -2
    # log.info('clearing-flag=%s' % pf)
    # if pf>1:
    #     return False
    # else:
    #     return True

    """ a version of hourly clearing """
    # if t[4]>55:
        # return False
    """ end of hourly clearing """
    if t[3] > 0 and t[3] < 23:
        return 3
    elif t[3] == 23 and t[4]<55:
        return 3
    elif t[3] == 2 and t[4]<5:
        return 1    # time shifting to a latter time than app call 0 play
    elif t[3] == 2 and (t[4]<10 and t[4]>5):
        return 2
    elif t[3] == 23 and t[4]>55:
        return 4
    else:
        return 3


@test("time = 1 @ 00:01:01")
def returntime1():
    t = time.struct_time((2011,1,1,0,0,1,1,1,1))
    assert window(t) == 1

@test("time = 2 @ 00:06:01")
def returntime2():
    t = time.struct_time((2011,1,1,0,6,6,1,1,1,1))
    assert window(t) == 2

@test("time = 3 @ 00:10:01 -- 25:55:00")
def returntime3():
    t1 = time.struct_time((2011,1,1,0,10,10,1,1,1,1))
    r1 = window(t1)
    t2 = time.struct_time((2011,1,1,1,1,1,1,1,1,1))
    r2 = window(t2)
    t3 = time.struct_time((2011,1,1,23,54,54,1,1,1,1))
    r3 = window(t3)
    assert r1 == 3 and r2 == 3 and r3 == 3

@test("time = 4 @ 23:55:00 - 0:0:0")
def returntime4():
    t = time.struct_time((2011,1,1,23,55,00,1,1,1,1))
    assert window(t) == 4
