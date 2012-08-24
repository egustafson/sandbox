#!/usr/bin/env python

import syslog

facility = syslog.LOG_USER
logOpt   = syslog.LOG_PID

def log(lvl, message):
    syslog.syslog(lvl, message)
    return

def info(message):
    log(syslog.LOG_INFO, message)
    return

syslog.openlog("test-syslog", logOpt, syslog.LOG_USER)

syslog.syslog("sending a test message, default level")

info("sending an info level message")

syslog.closelog()

