#!/usr/bin/env python

import gearman


print "Gearman Python Clent"
print " using: gearman.__version__: %s" % gearman.__version__
print " written for version:        2.0.2"

gm_client = gearman.GearmanClient(['localhost'])

req = gm_client.submit_job('demo', 'demo-data')
res = req.result

print "Result: %s" % res

print "done."

