# Package initializer
#
# In order for a directory to be evaluated as a package, (nested module),
# it must have a file named __init__.py located in it.
#
# ######################################################################

from PkgModule import MetaPublicClass

def pkgFunction():
    print "I am a package function, declared inside __init__.py."

print "I am initialization code inside of package/__init__.py."
