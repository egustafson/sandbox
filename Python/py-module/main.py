#!/usr/bin/env python

import Module

import package

from package.PkgModule import NestedClass


myHello = Module.Hello()
myHello.sayHello()


package.pkgFunction()

nestObj = NestedClass()
nestObj.hello()

# Note that this module was not explicitly imported, but
# rather used for a rename-import.
#
nestObj = package.PkgModule.OtherNestedClass()
nestObj.hello()


# The following class is rescoped by placing an
# "from package.PkgModule import Class" statement in
# the package's __init__.py file.  This would allow us
# to consolidate the public'esque classes and functions
# from nested modules into the parent package's namespace.
#
testObj = package.MetaPublicClass()
testObj.hello()

