Test Framework Reflection
=========================

This experiment will attempt to demonstrate using golang reflection to identify
test cases to run from within a set of registered objects.

Result:  this example seems to be proceeding in a positive direction.  It is
incomplete.  Currently there is no eloquent way to signal test failure and no
way at all to bail out of a test mid-test (i.e. `FailNow()`).  All of the
currently understood shortcomings have solutions.

2025-01-03 - Status:  I discovered a way to use the golang package `testing`
directly.  This is demo'ed in the [../testing-external] example.

In the future the strategy started in this demo may be suitable for use in a
distributed testing framework where each worker node comes up as a Blackboard
pattern allowing for tests to signal to multiple nodes that they need a set of
distributed participants in the test (or other workload) that is about to start.
