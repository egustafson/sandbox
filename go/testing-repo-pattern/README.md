Golang Package of Just Tests
============================

This example, and directory will show how a separate Golang repo, and
module, can be constructed with just tests, and possibly supporting
code for tests.

Directory:  testlib   # a Golang module with some demo test func's.
Directory:  testrepo  # a Golang module with "just tests"

Results:

1. The pattern appears to work just fine.
2. Spanning two golang modules inside of one Git repo makes a few
   things "look a little off", but the pattern never the less works.

Successful proof of concept.
