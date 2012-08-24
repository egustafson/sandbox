#!/usr/bin/env python

import sys
from omniORB import CORBA, PortableServer
import CosNaming, Example, Example__POA

class Echo_i (Example__POA.Echo):
    def echoString(self, mesg):
        print "echoString() called with message:", mesg
        return mesg

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)
poa = orb.resolve_initial_references("RootPOA")

ei = Echo_i()
eo = ei._this()

obj         = orb.resolve_initial_references("NameService")
rootContext = obj._narrow(CosNaming.NamingContext)

if rootContext is None:
    print "Failed to narrow to root naming context"
    sys.exit(1)

name = [CosNaming.NameComponent("test", "my_context")]
try:
    testContext = rootContext.bind_new_context(name)
    print "New test context bound"

except CosNaming.NamingContext.AlreadyBound, ex:
    print "Test context already exists"
    obj = rootContext.resolve(name)
    testContext = obj._narrow(CosNaming.NamingContext)
    if testContext is None:
        print "test.mycontext exists but is not a NamingContext"
        sys.exit(1)

name = [CosNaming.NameComponent("ExampleEcho", "Object")]
try:
    testContext.bind(name, eo)
    print "New ExampleEcho object bound"

except CosNaming.NamingContext.AlreadyBound:
    testContext.rebind(name, eo)
    print "ExampleEcho binding already existed -- rebound"

poaManager = poa._get_the_POAManager()
poaManager.activate()

orb.run()
