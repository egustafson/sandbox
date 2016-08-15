#!/usr/bin/env python

class Managed(object):

    def __init__(self, name):
        self._name = name

    def __enter__(self):
        print("__enter__")
        return self

    def __exit__(self, exc_type, exc_value, tb):
        print("__exit__")

    def prnt(self):
        print("I am [{}]".format(self._name))


if __name__ == "__main__":

    m = Managed("demo")

    with m as mg:
        mg.prnt()

    print("done.")
