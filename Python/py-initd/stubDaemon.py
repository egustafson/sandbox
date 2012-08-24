#!/usr/bin/env python


import os
import random
import signal
import sys
import time

_ShutdownFlag = False

def sighandler(signum, frame):
    global _ShutdownFlag
    print("received signal: ", signum)
    if signum == signal.SIGHUP:
        print("Reload - SIGHUP")
    if signum == signal.SIGTERM:
        _ShutdownFlag = True
        print("Shutdown - SIGTERM")

## ######################################################################

signal.signal(signal.SIGHUP, sighandler)  # Daemon reload signal
signal.signal(signal.SIGTERM, sighandler) # Daemon shutdown signal

pid = os.getpid()
print("[%d]: %s" % (pid, sys.argv))

while not _ShutdownFlag:
    signal.pause()

print("[%d] shutting down." % (pid))

sys.exit(0)
