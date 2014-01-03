#!/usr/bin/env python

import logging

FORMAT = '%(levelname)-8s [%(asctime)-15s] : %(message)s'
logging.basicConfig(format=FORMAT, level=logging.DEBUG)

if __name__ == '__main__':
    lvl_name = logging.getLevelName( logging.getLogger().getEffectiveLevel() )
    print("Logging at level: {}".format(lvl_name))

    logging.debug("debug message")
    logging.info("info message")
    logging.warn("warn message")
    logging.error("error message")
    logging.fatal("fatal message")
