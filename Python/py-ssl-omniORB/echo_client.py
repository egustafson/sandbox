#!/usr/bin/env python

import sys
import os
from omniORB import CORBA
import example
import omniORB.sslTP

omniORB.sslTP.certificate_authority_file('./CA/cacert.pem')
omniORB.sslTP.key_file('./CA/server.pem')
omniORB.sslTP.key_file_password('example')

# os.environ['ORBclientTransportRule'] = '* ssl'
os.environ['ORBtraceLevel'] = '25'

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)

# ior = sys.argv[1]
# obj = orb.string_to_object(ior)

ref = 'corbaloc:ssliop:thor.elfwerks:5678/EchoService'
# ref = 'corbaloc::thor.elfwerks:5678/EchoService'

print "Connecting to", ref

obj = orb.string_to_object(ref)

print "String to object complete."

eo = obj._narrow(example.EchoService)

print "Object narrowed."

if eo is None:
    print "Object reference is not an Example::Echo"
    sys.exit(1)

message = "Hello from Python"
result  = eo.echoString(message)

print "I said '%s'. The object said '%s'." % (message,result)
