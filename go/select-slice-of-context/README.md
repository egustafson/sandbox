Select on a Slice of Closing Channels
=====================================

i.e. select on `[]ctx.Done()`

Problem
-------

This example is born out of a need to monitor "from the side" a
variable set of channels and associated contexts.  The specific
use-case is a layer in the middle of a Watch functionality where
components can register, and de-register, watching for changes in a
subset of a repository (i.e. etcd like).

The "watch" is constructed from a context and an query (key) against
the repository and the resultant watch object is simply a Golang
channel.  In order to perform an orderly shutdown of the repository,
and consequently shutdown each watch, there must be a means of
canceling each context that was passed in on construction.  As the
repository, and it's watch() function, sit in a middle layer between
client and underlying store, the repository can wrap the caller's
context WithCancel(), holding on to the cancellation function and
passing the wrapped context down to the underlying store.  The
cancellation function can then be held, and used, by the repository
layer to cancel and tear-down each constructed watch at the time the
repository initiates destruction.

The catch, which will be demonstrated in this example is that each
watcher can be torn down (de-regitered) at an arbitrary time during
the life of the repository.  In order to prevent unbounded growth of a
list of watchers', context cancellation functions, the closure of the
watch channel will be indirectly detected by selecting against the
cancellation of the watcher's context, passed at construction.  In
order to detect cancellation, each active watcher's context must have
a select applied against it's Done() function.  In short, a select
against a slice of closing channels.

Solution(s)
-----------

The traditional approach of constructing a slice of channels and then
performing a select against the slice is afforded by Golang through
the `reflect` package and `reflect.Select()`.  Additionally, the
Golang community has posed an alternative approach which involves an
"aggregating channel" and individual goroutines per select.  Both
approaches are detailed in [stack-overflow-1][1].


<!-- References -->
[1]: <https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement> "how to listen to N channels? (dynamic select statement)"
