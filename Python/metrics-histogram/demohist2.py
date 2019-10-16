#!/usr/bin/env python3

import random

from appmetrics import metrics

if __name__ == '__main__':
    res = metrics.histogram.UniformReservoir(10288)
    hist = metrics.new_histogram("rand-spread", res)

    for rr in range(10000):
        v = random.random() * 1000
        hist.notify(v)

    print("Histogram:")
    for (k,v) in hist.get().items():
        print("  {}: {}".format(k,v))
    print("")
    for v in hist.get()['histogram']:
        print("  {}".format(v))
