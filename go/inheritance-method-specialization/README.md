Inheritance Method Specialization
=================================

This example shows a common (base) struct used to compose two separate, derived
classes that implement the same interface.  Each concrete class will implement
the same method, `Behavior()`, in a different way -- by returning a different
value.  

The base struct (class) does not implement the interface and would therefore be
considered an abstract class.