#!/usr/bin/env python
## ############################################################
##
##
## ############################################################

import os
import signal
import sys
import syslog
import threading
import time

## ############################################################

_SyslogIdent     = "py-initd"
_SyslogFacility  = syslog.LOG_DAEMON
_SyslogOpts      = syslog.LOG_PID
_SyslogUse       = True

_StartFile       = "daemons.cnf"

## ############################################################

ST_STOPPED  = 0
ST_STARTING = 1
ST_RUNNING  = 2

class ProcessHolder(threading.Thread):
    def __init__(self, ident, args):
        threading.Thread.__init__(self)
        self.ident    = ident
        self.cmd      = args[0]
        self.args     = args
        self.state    = ST_STOPPED
        self.shutdownFlag = False
        self.pid      = 0
        self.shortruns = 0


    def getStatus(self):
        stat = "[%d](%s): " % (self.pid, self.ident)
        if self.state == ST_RUNNING:
            for arg in self.args:
                stat += " %s" % (arg)
        elif self.state == ST_STOPPED:
            stat += "%s - STOPPED" % self.cmd
        elif self.state == ST_STARTING:
            stat += "%s - STARTING" % self.cmd
        else:
            stat += "%s - UNKNOWN STATE" % self.cmd
        return stat

    def getState(self):
        return self.state

    def getPid(self):
        return self.pid

    def shutdown(self):
        # Send a SIGTERM to the running PID, this DOES NOT ensure the PID is stopped.
        self.shutdownFlag = True
        if self.state == ST_STARTING:
            return False
        if self.state == ST_STOPPED:
            return True
        os.kill(self.pid, signal.SIGTERM)
        return True

        
    def run(self):
        while not self.shutdownFlag:
            self.state = ST_STARTING
            self.startTime = time.time()
            # Start (spawn) the process
            self.pid = os.spawnv(os.P_NOWAIT, self.cmd, self.args)
            self.state = ST_RUNNING
            log("started - %s" % (self.getStatus()))
            #
            # wait for the process to exit
            results = os.waitpid(self.pid, 0)
            #
            self.state = ST_STOPPED
            self.exitTime = time.time()
            if results[1] & 0xff > 0:
                resultString = "received signal(%d)" % (results[1] & 0x7f)
                if results[1] & 0x80 > 0:
                    resultString += "[core generated]"
            else:
                resultString = "exit code(%d)" % (results[1] >> 8)
            log("[%d](%s) terminated, %s" % (self.pid, self.ident, resultString))
            if self.startTime - self.exitTime < 2:
                self.shortruns += 1
            if self.shortruns > 5:
                log("(%s) - runaway daemon, terminating" % self.ident)
                self.shutdownFlag = True





## ############################################################

_ChildList = []
_PrStatus  = False
_Shutdown  = False

def sighandler(signum, frame):
    global _ChildList, _Shutdown, _PrStatus
    if signum == signal.SIGINT:
        log("Caught SIGINT, ignoring.  Please use SIGTERM.")
    if signum == signal.SIGHUP:
        _PrStatus = True
    if signum == signal.SIGTERM:
        _Shutdown = True
        while len(_ChildList) > 0:
            c = _ChildList.pop(0)
            if not c.shutdown():
                # Child state is ST_STARTING, wait for PID to be known
                # requeue at the back of the list for retry.
                _ChildList.append(c)

                
def sendStatus():
    global _ChildList
    for c in _ChildList:
        stat = c.getStatus()
        log(stat)


def log(msg, lvl=None):
    global _SyslogUse
    if _SyslogUse:
        if lvl:
            syslog.syslog(lvl, msg)
        else:
            syslog.syslog(msg)
    else:
        print(msg)
        
## ############################################################

if __name__ == '__main__':

    if _SyslogUse:
        syslog.openlog( _SyslogIdent, _SyslogOpts, _SyslogFacility )
        syslog.syslog("starting.")
    else:
        print("%s[%d] started." % (sys.argv[0], os.getpid()))

    signal.signal(signal.SIGHUP, sighandler)    # Daemon reload signal
    signal.signal(signal.SIGTERM, sighandler)   # Daemon shutdown signal
    signal.signal(signal.SIGINT, sighandler)    # ctrl-C (safeguard)

    daemonfile = file(_StartFile)
    for l in daemonfile:
        l = l.strip()
        if len(l) > 0 and l[0] != '#':
            try:
                (name, cmd) = l.split(":", 1)
                args = cmd.split()
                _ChildList.append( ProcessHolder(name, args) )
            except:
                log("ERROR: unable to parse - '%s'" % (l))

    for c in _ChildList:
        c.start()

    while not _Shutdown:
        if _PrStatus:
            sendStatus()
            _PrStatus = False
        signal.pause()
    
    for c in _ChildList:
        c.join()

    if _SyslogUse:
        syslog.syslog("exiting gracefully.")
    else:
        print "%s[%d] exiting gracefully." % (sys.argv[0], os.getpid())
    sys.exit(0)
