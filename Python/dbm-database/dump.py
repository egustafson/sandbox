#!/usr/bin/env python

import dbm.gnu

with dbm.gnu.open('demo.gdbm', 'r') as gdb:

    k = gdb.firstkey()
    while k is not None:
        print("{}".format(k.decode()))
        k = gdb.nextkey(k)

print("done.")
