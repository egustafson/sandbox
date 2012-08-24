#!/usr/bin/env python

f = open("raw.out", 'wb')

buf = ""
for ii in range(0,256):
    f.write(chr(ii))

f.close()
    
