#!/usr/bin/env python

import yaml

data = { 'a': 1, 'b': [2,3,4,5], 'c': {'x':'a', 'y':'b', 'z':'c'} }

print(yaml.dump(data))

print("done.")
