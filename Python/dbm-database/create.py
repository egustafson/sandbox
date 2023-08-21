#!/usr/bin/env python

from datetime import datetime

import dbm.gnu
import pickle

with dbm.gnu.open('demo.gdbm', 'cf',) as gdb:

    for n in range(1, 20):
        ip = "10.9.9.{}".format(n)
        gdb[ip] = "true"

    gdb.sync()

print("done.")
