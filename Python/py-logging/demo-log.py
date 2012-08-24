#!/usr/bin/env python2.5

import logging
import logging.config
import os

logConfigFile = 'log-config.ini'

if __name__ == '__main__':

    logfilename = "demo-%d.log" % (os.getpid())

    logging.config.fileConfig(logConfigFile, {'logfilename':logfilename} )

    logging.info("Logging Started")

    logging.debug("debug message")
    logging.info("info message")
    logging.warn("warn message")
    logging.error("error message")
    logging.fatal("fatal message")

    
