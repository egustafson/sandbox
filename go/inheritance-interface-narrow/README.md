Inheritance w/ Narrowing Interface
==================================

This example attempts to demonstrate two interfaces where the child interface
overrides a method in the parent to return a more narrowed type.

... Not directly possible.

And, it turns out that it does not appear possible to override a
method and narrow the method's signature in a child interface.

It is possible to "inherit" a parent interface (Animal in this
example) into a child interface (Cat) and in the child interface
define a _similar_, but named differently, method that does return the
narrowed type.  And then, simply define the wider method (Say()) to
invoke and return the narrower method and type (CatSay()).

I'm not sure that does a whole lot to improve readability
however. ... hummm.
