Golang Fn Node Streaming Framework
==================================

2nd attempt at a simple framework to hang functional (or functions)
off of and build streaming computation.

Concept:  compute a moving average, or variance (complexity theory)
of streaming data like stocks, or compute telemetry.

----

Framework philsophy:

* Construct a "Node" class that comes with the core, abstracted
  behaviors.
* Allow a single Function to be assigned to a "Node" that performs the
  specialization behavior of that Node -- i.e. compute an average.
  * Inputs are passed in as parameters.

* The core framework passes information through Channels.
* The generalized code manages the channels and invokes the
  specialized function with input, and captures output which is then
  passed to the output channel.   -- Plugin function doesn't handle
  channels.

* Each specialized function/Node runs as a spearate goroutine, but is
  provided as a simple 'func' that takes input in parameters and
  provides the results as the return value (struct-like) of the function.
