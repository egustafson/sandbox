#!/usr/bin/env python

import json

x = json.loads('{"foo":"var"}')
print x

x = {'key': 'value'}

print json.dumps(x)

