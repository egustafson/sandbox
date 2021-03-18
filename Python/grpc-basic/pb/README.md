gRPC IDL Files
==============

This directory hold the gPRC IDL ([Interface Definition
Language](https://en.wikipedia.org/wiki/Interface_description_language))
for this example. (`.proto` files)

Python
======

Due to what I consider limitations in Google's `protoc` IDL compiler
for gRPC/protobufs there is a bit of a hack that goes on in this
directory, combined with the examples Makefile (../Makefile).

_The Problem:_  It is desireable to place the generated code into a
Python Package directory -- under the premise that this demo could be
a shipable library.  The generated code is emited with what amounts to
a self referential renaming of it's own package.  In order to get the
importing correct, in Python 3, the path, or '.' needs to be included
in the packages name on the `import` line.  `protoc` for Python does
not support this; there are a number of GitHub issues open on this
topic that can be found by searching.  As best I can tell, at this
time (18 Mar 2021), there appears to be no concensus or general
solution being worked.  [this is an example, I didn't dig too deep.]

_The Solution:_ `protoc` will generate the desired import path **IF**
the location of the `.proto` file is in just the right path, combined
with the "right" flags.  The path for the `.protoc` file would be both
confusing and problematic to other languages including the same file
-- one of IDL's, and specifically gRPC's, strengths is cross language
use.  The 'hack' is that the `Makefile`, combined with `.gitignore`
create and then don't include in the checked in files, the `.prococ`
file in the proper path, for Python.  You will note that a sub-directory
named 'demo' will contain a symlink back to this directory
`demo.proto` file -- this is the "hack".
