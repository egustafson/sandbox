#!/usr/bin/env python

import sys
from omniORB import CORBA
import CosNaming


servant_name = "examples/servant"

# Example Client

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)

nsobj   = orb.resolve_initial_references("NameService")
rootExt = nsobj._narrow(CosNaming.NamingContextExt)

if rootExt is None:
    print "Failed to narrow the root naming context - exiting."
    sys.exit(1)

exobj = rootExt.resolve_str(servant_name)

import ex

exsvnt = exobj._narrow(ex.servant)

exsvnt.ping()
exsvnt.shutdown()

# End of client
