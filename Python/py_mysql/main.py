#!/usr/pkg/bin/python

import MySQLdb

connection = MySQLdb.connect(host="localhost", user="demo", db="test")
cursor = connection.cursor()
cursor.execute("select * from testac")

data = cursor.fetchall()
fields = cursor.description

for des in fields:
    print "Field: ", des

print "----"

for row in data:
    print "Row: ", row
