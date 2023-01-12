Managed Worker Pool
===================

Goal:  Implement a worker pool that allows the pool to interact with
each worker both upon return to the pool and just prior to issuance as
a worker.  Additionally, the pool should be able to scale the pool up
and down based on both high/low water marks and based on external
changes to those levels.
