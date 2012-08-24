#!/usr/bin/python

import gc


class Gar:
    def __init__(self, theText):
        self.myText = theText
        self.val    = 1


queue = []
index = 0
while index < 100:
    index = index + 1;
    queue.append(Gar("Object #" + str(index)))
    gc.collect()
    print "GC object count:  " + str(len(gc.get_objects()))

print "Queue has " + str(len(queue)) + " elements."

queue = []
gc.collect()
print "GC object count:  " + str(len(gc.get_objects()))
