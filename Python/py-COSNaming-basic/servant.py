#!/usr/bin/env python

import sys
from omniORB import CORBA
import CosNaming

servant_name = "examples/servant"

# Example Servant

import ex, ex__POA

class Servant_impl(ex__POA.servant):
    def __init__(self, orb):
        self.orb = orb
        return
    
    def ping(self):
        print "received ping."
        return

    def shutdown(self):
        print "received shutdown."
        try:
            rootCtx = self.orb.resolve_initial_references("NameService")._narrow(CosNaming.NamingContextExt)
            unbindObjToName(servant_name, rootCtx)
        except:
            print "Failed to unbind the servering from COSNaming"
        try:
            self.orb.shutdown(0)
        except:
            print "Failed to send shutdown signal to the ORB"
        return


# ############################################################

def unbindObjToName(namestr, rootCtx):
    objName = rootCtx.to_name(namestr)
    try:
        rootCtx.unbind(objName)
    except:
        print "Failed to unbind name:  %s" % namestr
    return

def bindObjToName(obj, namestr, rootCtx):
    ctxName = rootCtx.to_name(namestr)
    objName = ctxName.pop()

    boundName = []
    for n in ctxName:
        boundName.append(n)
        print "trying:  %s.%s" % (n.id, n.kind)
        try:
            rootCtx.bind_new_context(boundName)
            print "bound to naming context"
        except CosNaming.NamingContext.AlreadyBound:
            ctx = rootCtx.resolve(boundName)._narrow(CosNaming.NamingContext)
            print "found bound context, continuing"
            if ctx is None:
                print "context for servant could not be created or resolved - exiting."
                sys.exit(1)

    ctxName.append(objName)
    try:
        print "binding object:  %s" % rootCtx.to_string(ctxName)
        rootCtx.bind(ctxName, obj)
    except CosNaming.NamingContext.AlreadyBound:
        rootCtx.rebind(ctxName, obj)
    return


# ############################################################

if __name__ == "__main__":
    
    orb = CORBA.ORB_init(sys.argv, CORBA.ORB_ID)
    poa = orb.resolve_initial_references("RootPOA")

    exref  = Servant_impl(orb)._this()

    obj         = orb.resolve_initial_references("NameService")
    rootContextExt = obj._narrow(CosNaming.NamingContextExt)

    if rootContextExt is None:
        print "Failed to narrow the root, ext naming context"
        sys.exit(1)

    bindObjToName(exref, servant_name, rootContextExt)

    poaManager = poa._get_the_POAManager()
    poaManager.activate()

    print "Running..."
    orb.run()
    print "Quit ORB, exiting."

# end of servant.py
