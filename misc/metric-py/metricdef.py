#!/usr/bin/env python

import collections
import copy
import hashlib
import json
import time

class Metric:
    def __init__(self, md, value, ts=None):
        self.md    = md
        self.value = value
        if ts:
            self.ts = ts
        else:
            self.ts = time.time()

    def to_json(self, pretty=False, indent=False):
        m = copy.deepcopy(self.md.m)
        m['value'] = self.value
        m['ts'] = self.ts
        sep = (",", ":")
        if pretty:
            sep = (", ", ": ")
        ind = None
        if indent:
            ind = 2
        return json.dumps(m, separators=sep, indent=ind)
        

class MetricDef:
    def __init__(self, **kwargs):
        self.hk = None # invalidate the metric def's hashKey
        self.m = {}
        self.m['coords'] = {}
        if 'ns' in kwargs:
            self.m['ns'] = kwargs['ns']
            del kwargs['ns']
        if 'coords' in kwargs:
            self.m['coords'] = kwargs['coords']
            del kwargs['coords']
        if len(kwargs) > 0:
            self.addCoords(**kwargs)

    def setNamespace(self, ns):
        self.hk = None # invalidate the metric def's hashKey
        self.m['ns'] = ns

    def getNamespace(self):
        return self.m['ns']

    def setCoords(self, **kwargs):
        self.hk = None # invalidate the metric def's hashKey
        self.m['coords'] = {}
        self.addCoords(**kwargs)

    def getCoords(self):
        return self.m['coords']

    def addCoord(self, k, v):
        self.hk = None # invalidate the metric def's hashKey
        self.m['coords'][k] = v

    def hasCoord(self, k, v=None):
        return k in self.m['coords'] and ( v == None or self['coords'][k] == v )

    def addCoords(self, **kwargs):
        self.hk = None # invalidate the metric def's hashKey
        for (k,v) in kwargs.items():
            self.m['coords'][k] = v

    def rmCoord(self, k):
        self.hk = None # invalidate the metric def's hashKey
        if k in self.m['coord']:
            del self.m['coord'][k]
    
    def rmCoords(self, *keys):
        self.hk = None # invalidate the metric def's hashKey
        for k in keys:
            self.rmCoord(k)

    def rmCoords(self, **kwargs):
        self.hk = None # invalidate the metric def's hashKey
        for k in kwarg.iteritems():
            self.rmCoord(k)

    def hasCoords(self, **kwargs):
        for (k,v) in kwargs.items():
            if not self.hasCoord(k,v):
                return False
        return True

    def hasCoords(self, *keys):
        for k in keys:
            if not self.hasCoord(k,v):
                return False
        return True

    def to_json(self, pretty=False, indent=False):
        sep = (",", ":")
        if pretty:
            sep = (", ", ": ")
        ind = None
        if indent:
            ind = 2
        return json.dumps(self.m, separators=sep, indent=ind)

    def hashKey(self):
        if not self.hk:
            sep = (",", ":")
            material = json.dumps(self.m, separators=sep, sort_keys=True)
            self.hk = hashlib.md5(material).hexdigest()
        return self.hk
    
    def emit(self, value, ts=None):
        return Metric(self, value, ts)


class MetricDefTable:

    def __init__(self):
        self.nsdict = collections.defaultdict(set)
        self.coords = collections.defaultdict(set)

    def add(self, md):
        self.nsdict[md.getNamespace()].add(md)
        for (k,v) in md.getCoords().iteritems():
            self.coords["{}:{}".format(k,v)].add(md)

    def find(self, ns, **kwargs):
        if ns not in self.nsdict:
            return set()
        matches = self.nsdict[ns]
        for (k,v) in kwargs.iteritems():
            kv = "{}:{}".format(k,v)
            if kv not in self.coords: 
                return set()
            matches.intersection_update(self.coords[kv])
        return matches

    def allEntries(self):
        entries = set()
        for s in self.nsdict.itervalues():
            for md in s:
                entries.add(md)
        return entries


if __name__ == '__main__':

    mdTable = MetricDefTable()

    md = MetricDef()
    md.setNamespace('test.namespace')
    md.addCoord('id', 'xyzzy')
    md.addCoord('val', 'abc')

    mdTable.add(md)

    print(md.to_json(True))
    print(md.hashKey())

    m = md.emit(123.45)
    print(m.to_json(True))

    md2 = MetricDef(ns='test.ns2', t='asdfg', v='12345')
    mdTable.add(md2)
    print(md2.to_json(True))

    mdTable.add( MetricDef(ns='test.ns2', t='asdfg', v='1') )
    mdTable.add( MetricDef(ns='test.ns2', t='asdfg', v='2') )
    mdTable.add( MetricDef(ns='test.ns2', t='asdfg', v='3') )
    mdTable.add( MetricDef(ns='test.ns2', t='asdfg', v='4') )
    mdTable.add( MetricDef(ns='test.ns2', t='xyz', v='1') )
    mdTable.add( MetricDef(ns='test.ns2', t='xyz', v='2') )
    mdTable.add( MetricDef(ns='test.ns2', t='xyz', v='3') )
    mdTable.add( MetricDef(ns='test.ns2', t='xyz', v='4') )

    print("-- All metric definitions")
    for md in mdTable.allEntries():
        print("{} : {}".format(md.hashKey(), md.to_json(True)))

    print("-- All ns=test.ns2 + t=asdfg")
    for md in mdTable.find('test.ns2', t='asdfg'):
        print(md.to_json(True))

    print("done.")
