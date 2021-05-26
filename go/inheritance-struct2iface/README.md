Golang Inheritance with Method Override
---------------------------------------

This example shows how **Method Override** can be implemented in
Golang.

Technically this is not inheritance, but it has the feeling of it.
There is a "child" struct (or interface) that implements a method,
`Close()`.  And then the child is included as an aggregate of the
parent.  IF the parent does not implement the child's method, see
`Op()`, then it directly callable by users of the parent struct.  IF
the parent implements the same method as the child, `Close()` then the
parent's implementation overrides the child.  Finally, the parent's
"overriding" method can explicitly invoke the child's method as is
done in `Parent.Close()`.  (It is also possible for users of the
parent to explicitly invoke the child's (overridden) method).
