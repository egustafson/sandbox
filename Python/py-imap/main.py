#!/usr/bin/env python

import getpass, imaplib

m = imaplib.IMAP4('spot.elfwerks.org')

uname = "ericg@gustafson-consulting.com"
pword = getpass.getpass()

print "User: %s, Pass: %s" % (uname, pword)

m.login(uname, pword)

folders = m.list()

print "folder count:", len(folders)
for f in folders[1]:
    print f

# m.select()
# typ, data = m.search(None, 'ALL')
# for num in data[0].split():
#     typ, data = m.fetch(num, '(RFC822)')
#     print 'Message %s\n%s\n' % (num, data[0][1])

# m.close()
m.logout()
