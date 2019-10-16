#!/usr/bin/env python3

import random
import time

from appmetrics import metrics


def do_oper():
    if random.random() < 0.1:
        # occasionally (10%) introduce a large variance
        time.sleep(0.1)

    for ii in range(10):
        time.sleep(0.00001)  # 10 uS


if __name__ == "__main__":

    hist = metrics.new_histogram("call_time")

    loops = 100
    for ii in range(loops):
        start_ns = time.time_ns()
        do_oper()
        end_ns = time.time_ns()
        hist.notify(end_ns - start_ns)

    print("Histogram: \n")
    for (k, v) in hist.get().items():
        print("  {}: {}".format(k, v))
