#!/usr/bin/env python

import shelve

from datetime import datetime


PFILENAME = 'demo_persistence.dbm'

with shelve.open(PFILENAME) as db:

    for n in range(2, 20):
        ip = "10.8.8.{}".format(n)
        rec = {
            'ip': ip,
            'last_seen': datetime.utcnow(),
        }
        db[ip] = rec

print('done.')

