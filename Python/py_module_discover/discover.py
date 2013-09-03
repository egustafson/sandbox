#!/usr/bin/env python

import os
import sys

print("sys.path:")
for p in sys.path:
    print("  {}".format(p))

if 'PYTHONPATH' in os.environ:
    python_path = os.environ['PYTHONPATH']
    print("PYTHONPATH:  {}".format(python_path))

sys.path.append("./discover_dir")
import simple
import proto

print("done.")

