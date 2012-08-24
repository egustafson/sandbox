#!/usr/bin/env python

import xml.dom.ext
from xml.dom.ext.reader import Sax2
from xml import xpath

filename = "example.xml"

print "Demo started."

reader = Sax2.Reader()
doc = reader.fromStream(filename)

#xml.dom.ext.PrettyPrint(doc)
#xml.dom.ext.Print(doc)

mod_list = xpath.Evaluate('/configuration/module', doc.documentElement)

for mod in mod_list:
    name = xpath.Evaluate("@name", mod)[0].nodeValue
    print "Module Name:  %s" % name
    param_list = xpath.Evaluate('parameter', mod)
    for param in param_list:
        name = xpath.Evaluate('@name', param)[0].nodeValue
        val  = xpath.Evaluate('text()', param)[0].nodeValue
        print "%s = %s" % (name, val)
        
