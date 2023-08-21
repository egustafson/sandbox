#!/usr/bin/env python

import shelve

from datetime import datetime


PFILENAME = 'demo_persistence.dbm'

with shelve.open(PFILENAME) as db:

    for k in iter(db.keys()):
        rec = db[k]
        print("{}: {}z".format(rec['ip'], rec['last_seen'].isoformat(timespec='seconds')))

print("done.")
