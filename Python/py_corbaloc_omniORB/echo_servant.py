#!/usr/bin/env python

import sys
import os
from omniORB import CORBA, PortableServer
import Example, Example__POA

class Echo_i (Example__POA.Echo):
    def echoString(self, mesg):
        print "echoString() called with message:", mesg
        return mesg


# ===== Main =====

os.environ['ORBendPoint'] = 'giop:tcp::5678';

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)
ins_poa = orb.resolve_initial_references("omniINSPOA")

ei = Echo_i()
# eo = ei._this()

eo = ins_poa.activate_object_with_id("ExampleEcho", ei)

# For some reason, the reference returned does not print properly.
print orb.object_to_string(eo)

poaManager = ins_poa._get_the_POAManager()
poaManager.activate()

orb.run()
