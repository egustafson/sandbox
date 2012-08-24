#!/usr/pkg/bin/python

import MySQLdb

connection = MySQLdb.connect(host="localhost", user="demo", db="test")
cursor = connection.cursor()
cursor.execute("select * from testac")

data = cursor.fetchall()
fields = cursor.description

print ""
print "paramstyle = ", MySQLdb.paramstyle
print ""

for des in fields:
    print "Field: ", des

print "----"
print "Retreived", cursor.rowcount, "rows."

for row in data:
    print "Row: ", row
