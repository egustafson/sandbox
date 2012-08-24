#!/usr/bin/env python

import sys
from omniORB import CORBA
import CosNaming

# Example

orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)

obj         = orb.resolve_initial_references("NameService")
print "NameService ior:", orb.object_to_string(obj)
rootContext = obj._narrow(CosNaming.NamingContext)

if rootContext is None:
    print "Failed to narrow the root naming context"
    sys.exit(1)

print "Narrowed naming context - COSNaming present."


rootContextExt = obj._narrow(CosNaming.NamingContextExt)

if rootContextExt is None:
    print "Failed to narrow the root, ext naming context"
    sys.exit(1)

ctxPath = rootContextExt.to_name("a/b/c")

print "Context Path:"
for n in ctxPath:
    if n.kind:
        print "  %s.%s" % (n.id, n.kind)
    else:
        print "  %s" % (n.id)


