#!/usr/bin/env python

import gearman


def gm_task(gm_worker, job):
    print "working..."
    print "  job.handle: %s" % job.handle
    print "  job.task:   %s" % job.task
    print "  job.unique: %s" % job.unique
    print "  job.data:   %s" % job.data
    return "response: %s" % job.data




print "Gearman Python Worker"
print " using: gearman.__version__: %s" % gearman.__version__
print " written for version:        2.0.2"
print ""
print "started ..."

worker = gearman.GearmanWorker(['localhost'])
worker.register_task('demo', gm_task)
worker.work()

print "done."

