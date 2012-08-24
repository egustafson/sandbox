#!/usr/bin/env python

import threading, Queue

class ConsumerThread(threading.Thread):
    
    def __init__(self, workq):
        threading.Thread.__init__(self)
        self.wq = workq

    def run(self):
        print "Consumer Running"
        while 1:
            work_item = self.wq.get()
            print "Consumed: ", work_item
        return



## Main

print "Multi Threaded Queue Example"

workQueue = Queue.Queue()

consumer = ConsumerThread( workQueue )
consumer.start()

workQueue.put("Item 1")
workQueue.put("Item 2")
