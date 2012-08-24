#!/usr/bin/env python


import os
import random
import sys
import time


pid = os.getpid()
sleepTime = random.randint(1,5)
print "PID(%d) starting, sleeping for %d seconds." % (pid, sleepTime)
time.sleep(sleepTime)
print "PID(%d) exiting." % (pid)

#sys.exit(1)
