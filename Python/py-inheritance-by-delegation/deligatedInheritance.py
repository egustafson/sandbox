#!/usr/bin/env python
## ################################################################################
##
##   This example shows how to create the effect of a derived class (Inheritance)
## when you can't use explicit inheritance in Python.  Its kind of like adding
## a method to an object after the object has been created.
##
##   Problem - we want to wrap a MySQLdb connection object and add a destructor
## that either returns the connection to a pool, or closes the connection when
## it goes out of scope.  { Ensuring resources don't leak. }
##   In this case the object is created by a library for us (sqlobject) which
## means we can't extend the class because we don't have access at creation time.
##   Solution - Use delegation instead of inheritance.
##
## ################################################################################

class Base:
    def __init__(self):
        print "object of type Base initialized."
        return

    def baseMethod(self):
        print "baseMethod invoked."
        return


class Delegate:
    def __init__(self, base):
        print "object of type Delegate initialized."
        # save a reference to 'base'
        self.__dict__['base'] = base
        return

    def delegateMethod(self):
        print "delegateMethod invoked."
        return

    def __del__(self):  # Our extension - the destructor
        print "destructor for Delegate class fired."
        return

    # Pass everything else through to the Base by
    # using __{get,set}attr__  (deligated inheritance)
    #
    def __getattr__(self, attr):
        return getattr(self.base, attr)
    def __setattr__(self, attr, value):
        return setattr(self.base, attr, value)

#
# Example usage
#
if __name__ == '__main__':
    b = Base()   # create our Base object (before the delegate)
    #
    d = Delegate(b)  # initialze Delegate with the _already created_ base
    #
    d.delegateMethod()
    d.baseMethod()
    #
    # destructor will fire as script terminates
    #
    print "done."
