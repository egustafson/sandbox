Main Program Testing w/ golang test framework
=============================================

The goal of this example is to build a "main program" that when run can execute
a set of test cases and to use golang's testing package along with (likely) the
stretchr/testify packages.

Use-case:  build a stand alone testing tool that fits the space of integration
testing and can be incrementally added to using the same patterns as golang uses
for tests.