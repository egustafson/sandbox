#!/usr/bin/env python

import sys
from omniORB import CORBA
import CosNaming, Example

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)

obj         = orb.resolve_initial_references("NameService")
rootContext = obj._narrow(CosNaming.NamingContext)

if rootContext is None:
    print "Failed to narrow the root naming context"
    sys.exit(1)

name = [CosNaming.NameComponent("test", "my_context"),
        CosNaming.NameComponent("ExampleEcho", "Object")]
try:
    obj = rootContext.resolve(name)

except CosNaming.NamingContext.NotFound, ex:
    print "Name not found"
    sys.exit(1)

eo = obj._narrow(Example.Echo)

if eo is None:
    print "Object reference is not an Example::Echo"
    sys.exit(1)

message = "Hello from Python"
result  = eo.echoString(message)

print "I said '%s'. The object said '%s'." % (message,result)
