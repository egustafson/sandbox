System Test Cases with pytest Fixtures
======================================

This example attempts to demonstrate how a "System Test" case (or
suite) can be implemented with the assistance of pytest's fixtures and
mark(s).

The goal:

1. Have an "external", running service that could have been started
   locally (a CI environment)
2. Run "system test cases" that interact with the "service under test"
3. NOT Run the "system test cases" during normal unit-testing,
   (i.e. simply running `pytest` at in the directory would not cause
   the system test cases to run)
