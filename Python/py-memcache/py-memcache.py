#!/usr/bin/env python

import memcache

mc = memcache.Client(['127.0.0.1:11211'], debug=0)

mc.set("key", "value")
value = mc.get("key")

print "value = '" + value + "'"

mc.delete("key")

print "done."
