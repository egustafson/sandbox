#!/usr/bin/env python

import socket

#RRDCACHED_ADDRESS = "/opt/collectd/var/run/collectd-unixsock"
RRDCACHED_ADDRESS = "/var/run/rrdcached.sock"

print "Opening: ", repr(RRDCACHED_ADDRESS)

s = socket.socket(socket.AF_UNIX)
s.connect(RRDCACHED_ADDRESS)
s.send("FLUSHALL\n")
data = s.recv(4096)
s.close()
print "Response: \n", repr(data)

print "\ndone.\n"
