#!/usr/bin/env python

import unittest
import sys

VERBOSE = False

class ExampleUnitTest(unittest.TestCase):

    def setUp(self):
        if VERBOSE: print("setUp")

    def tearDown(self):
        if VERBOSE: print("tearDown")

    def test_one(self):
        if VERBOSE: print("test_one")

    def test_two(self):
        if VERBOSE: print("test_two")


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == '-v': VERBOSE = True
    unittest.main()
