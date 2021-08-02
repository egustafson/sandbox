# -*- utf-8 -*-

# A demo program.

import time

from threading import Thread, Event


class Ticker(Thread):

    def __init__(self, fn, period=1):
        super().__init__()
        self.period = period
        self.fn = fn
        self.ev = Event()

    def run(self):
        print("-> bg starting")
        while not self.ev.wait(self.period):
            self.fn()
        print("-> bg exiting - finalized")

    def stop(self):
        self.ev.set()


class Worker:

    def __init__(self):
        self._ticker = Ticker(fn=self.bg_work, period=1)
        self._ticker.start()

    def bg_work(self):
        print("-> bg work")

    def fg_work(self):
        print("fg work")

    @property
    def ticker(self):
        return self._ticker


if __name__ == '__main__':
    print('Main')
    w = Worker()
    for ctr in range(4):
        time.sleep(1.3)
        w.fg_work()
    print("finalizing worker")
    w.ticker.stop()
    for ctr in range(4):
        time.sleep(1.3)
        w.fg_work()
    print("done.")
