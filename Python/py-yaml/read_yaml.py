#!/usr/bin/env python

import yaml

with open("creds.yaml") as f:
    data = yaml.load(f)
    # print(data)

creds = data['creds']
for (k,v) in creds.items():
    print("{} : {}".format(k,v))

print("done.")
