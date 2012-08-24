#!/usr/bin/python

import time

ii = 0
lapTime = time.time()
tick = 3.0

while ii < 300:
    ii = ii + 1
    if (lapTime + tick) < time.time():
        print time.asctime() + " : ii = " + str(ii)
        lapTime = time.time()

    time.sleep(1)
