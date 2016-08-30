#!/usr/bin/env python

## Write to the systemd journal
##
## watch the journal in another window with:
##
##   sudo journalctl -f  {--output=json}
##

from systemd import journal

# simple message send
#
journal.send("a systemd-journal log message from python land.")

# send message w/ extra fields.
#
journal.send( "A systemd-journal log message from python land, with extra field",
              EXTRAFIELD='foo field')

# sendv() is also supported
#
journal.sendv( 'MESSAGE=systemd-journal log message using sendv() from python',
               'EXTRAFIELD=foo field' )

# a very syslog like message
#
journal.send( 'Principal log message (from python land)',
              SYSLOG_FACILITY='USER',
              SYSLOG_IDENTIFIER='py-write-journal',
              SYSLOG_PID=-1,
              PRIORITY='NOTICE' )

## Local Variables:
## mode: python
## End:
