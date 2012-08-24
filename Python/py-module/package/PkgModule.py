#

class NestedClass:
    def hello(self):
        print "Hello from NestedClass inside package.PkgModule."


class OtherNestedClass:
    def hello(self):
        print "Hello from OtherNestedClass inside package.PkgModule."


class MetaPublicClass:
    """This class is intended to be imported and rescoped into the parent package"""
    def hello(self):
        print "Hello from MetaPublicClass inside package.PkgModule."

print "package.PkgModule initialized."

