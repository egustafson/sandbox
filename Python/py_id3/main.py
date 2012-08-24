#!/usr/pkg/bin/python

from ID3 import *
try:
    filename = './DontGo.mp3'
    id3info = ID3(filename)

    print id3info
    
except InvalidTagError, message:
    print "Invalid ID3 tag:", message

