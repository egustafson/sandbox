#!/usr/bin/env python

import sys
from omniORB import CORBA, PortableServer
import Example, Example__POA

class Echo_i (Example__POA.Echo):
    def echoString(self, mesg):
        print "echoString() called with message:", mesg
        return mesg

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)
poa = orb.resolve_initial_references("RootPOA")

ei = Echo_i()
eo = ei._this()

poaManager = poa._get_the_POAManager()
poaManager.activate()

message = "Hello"
result  = eo.echoString(message)

print "I said '%s'. The object said '%s'." % (message,result)
