#!/usr/bin/env python

import sys
import os
from omniORB import CORBA, PortableServer
import omniORB.sslTP

import Example, Example__POA

class Echo_i (Example__POA.Echo):
    def echoString(self, mesg):
        print "echoString() called with message:", mesg
        return mesg


# ===== Main =====

os.environ['ORBendPoint'] = 'giop:ssl::5678'

omniORB.sslTP.certificate_authority_file('./CA/cacert.pem')
omniORB.sslTP.key_file('./CA/server.pem')
omniORB.sslTP.key_file_password('example')


orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)
ins_poa = orb.resolve_initial_references("omniINSPOA")
# poa = orb.resolve_initial_references("RootPOA")

ei = Echo_i()

# For some reason, the reference returned does not print properly.
# eo = ei._this()
eo = ins_poa.activate_object_with_id("ExampleEcho", ei)

print orb.object_to_string(eo)

poaManager = ins_poa._get_the_POAManager()
# poaManager = poa._get_the_POAManager()
poaManager.activate()

orb.run()
