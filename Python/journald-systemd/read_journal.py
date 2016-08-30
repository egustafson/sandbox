#!/usr/bin/env python

## Read and output the systemd journal.
##
##   start:  when the system booted (i.e. 'this_boot()')
##   end:    most recent log entry.
##
## run as root to see the system events
##

import json

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
j.this_boot()          ## only consume logs since the most recent reboot

#
# utilize the built in Iterator pattern
#
for entry in j:
    rec = json.dumps(entry, separators=(',',':'))
    print(rec)

# prints all logs since boot and then exits.

## Local Variables:
## mode: python
## End:
