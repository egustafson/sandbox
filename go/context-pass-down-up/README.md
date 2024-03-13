Context Value - passing back from initialization function
=========================================================

Demonstrates how to encode values onto a context.Context through a sequential
chain of function calls.  This method may be useful during a high level
initialization of components such as start-up.

Problem:  in order to attach a value to a Context the WithValue() method is used
and this method acomplishes the attachment by returning a new Context.  The
underlying implementation of the context does not support adding values, but
rather creates a new Context with the value added.

Solution: follow the same pattern as context.WithValue and have each function in
the chain return the (possibly updated) context as a return value.


Initial Question
----------------

Q:  Is the context.Context passed into a function by reference?

Reason:  If the context is passed in by reference then it is possible to attach
a value to the context during an initialization function and have that value
persist upon return, and into subsiquent initialization functions.

Is this a good practice?  Unclear, this example is merely attempting to validate
the premise of the question.