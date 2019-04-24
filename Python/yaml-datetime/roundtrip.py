#!/usr/bin/env python

import yaml

with open("input.yml") as f:
    data = yaml.load(f)

print("deserialized: {!r}".format(data))

print("serialized: \n")
print(yaml.dump(data))
print("")
