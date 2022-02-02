Pipelined Work with Pooled Workers
----------------------------------

This demo constructs a pattern where there's a queue of work (a
buffered chan), a pool of workers (a buffered chan), and a dispatcher
in the middle.

The pool of workers are refreshed after "stale_worker_time".  This
emulates the notion of a pool of connectors to a (network) service
where the connectors are cycled after some period.

The Dispatcher sits in the middle and performs the following
algorithm:

```
1. Block waiting for worker
2. If worker > maxAge then Refresh(worker) and start loop again.
3. Block waiting for either:
    a. work in from the work queue
    b. timeout (short time ~1s)
4. if:
    a. work -> async dispatch worker(work)
    b. timeout -> return worker to back of queue
5. loop to beginning
```
