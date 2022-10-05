Struct TYPE as a Reference Id
=============================

Q: can the Type of a Golang object used as a unique identifier.

Reason: If the type of an object can be used as a unique key then
objects, like a context object, that hold objects of uniquely
different types can use the type itself as the index to retrieve the
object.

Conclusion:

- the nil value of a type "StructType{}" acts as an ID.
- aliased types === original type with respect to the type's "value".
- extended types are unique versus the type extended from.

Using the syntax "StructType{}" in multiple locations in a program
appears to yield the same value and can be used as a unique key in a
map.
