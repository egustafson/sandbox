#!/usr/bin/env python

import sys
import CORBA
import IDL.mod

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)

print "The IDL string is: %s" % (IDL.mod.s)
