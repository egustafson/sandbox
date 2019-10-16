#!/usr/bin/env python3

import time

## Warning:  requires Python 3.7+
##   uses time.time_ns() introduced in 3.7
##

def loop_time_ns(count):
    begin_ns = time.time_ns()
    for ii in range(count):
        t = time.time_ns()
    end_ns = time.time_ns()
    return (end_ns - begin_ns)

def loop_modulo(count):
    begin_ns = time.time_ns()
    for ii in range(count):
        x = ii % 111
    end_ns = time.time_ns()
    return (end_ns - begin_ns)


if __name__ == '__main__':

    rr = 1000000

    duration_ns = loop_time_ns(rr)
    dur_per = duration_ns / rr
    print("Time for {0} calls to time.time_ns():  {1}ns ({2}ns per call)".format(rr, duration_ns, dur_per))

    duration_ns = loop_modulo(rr)
    dur_per = duration_ns / rr
    print("Time for {0} calls to modulo():  {1}ns ({2}ns per call)".format(rr, duration_ns, dur_per))

    print("Done.")
