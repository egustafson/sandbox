#!/usr/bin/env python

## Poll systemd journal and output (i.e. tail)
##
##   start:  the next event to arrive (i.e. seek_tail())
##   end:    -- never -- (use ^C to exit)
##
## run as root to see the system events
##

import json
import select

from systemd import journal


## JsonReader modifies the standard journal.Reader class
## so that the entries returned by the _Python_ journal.Reader
## can be (trivially) converted into JSON.
##
class JsonReader(journal.Reader):

    def _convert_field(self, key, value):
        if isinstance(value, (list, tuple)):
            return value[0]
        else:
            return value


#j = journal.Reader()  ## read _everything_ (presume running as root)
j = JsonReader()
j.seek_tail()          ## advance to the end of the journal.
j.get_next()           ## make sure we've (pre)consumed everything.

p = select.poll()
p.register(j, j.get_events())
while True:
    p.poll()
    for entry in j:
        rec = json.dumps(entry, separators=(',',':'))
        print(rec)

## loops forever, use ^C to exit.

## Local Variables:
## mode: python
## End:
