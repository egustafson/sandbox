#!/usr/bin/env python
## ########################################
"""
An example program to demonstrate detecting open/closed stderr.
"""
## ########################################
## $Id$

import sys
import os

## ########################################

## ########################################
##      Main Program
##
if __name__ == '__main__':

    print "Controlling Terminal:  %s" % os.ctermid()

    print "stdout isatty():  %r" % sys.stdout.isatty()
    print "stdout fd:        %d" % sys.stdout.fileno()
    print "stdout name:      %s" % sys.stdout.name
    print "stdout closed:    %r" % sys.stdout.closed

    stdoutfd = sys.stdout.fileno()
    os.close(stdoutfd)
    print >> sys.stderr, "stdout closed:    %r" % sys.stdout.closed
    print >> sys.stderr, "fstat(2):  %r" % os.fstat(stdoutfd)
    os.write(stdoutfd, "bogus")
    print "string to stdout after calling os.close()"

    sys.stdout.close()
    ##print "String to stdout"
    print >> sys.stderr, "stdout closed:    %r" % sys.stdout.closed



    stderr = sys.stderr
    try:
        print "Before setting sys.stdout to None"
        print >> sys.stderr, "Sent to stderr"
        sys.stdout = None
        sys.stderr = None
        print "After setting sys.stdout to None"
    except:
        print >> stderr, "Caught exception"
