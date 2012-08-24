#!/usr/bin/python

import gc

startObjList = gc.get_objects()
ii = 0
jj = 1
endObjList = gc.get_objects()

gc.set_debug(gc.DEBUG_STATS|gc.DEBUG_LEAK)

gc.collect()

n = (jj-1) + ii
for x in endObjList:
    for y in startObjList:
        if x is y:
            break
    else:
        print "Found one:  %d" % n
    n = n + 1

n = 542
print "About to view object index " + str(n)
print endObjList[n]

ii = 42
jj = ii + 21
ii = jj - 32


gc.collect()
