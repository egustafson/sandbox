#!/usr/bin/env python
## ############################################################
##   String 'variable' substutition in Python w/ RE's.
##
##   This example is ment to show a routine that detects and
## then substitutes from a dictionary 'variable' expressions
## that are embedded in a string.
##
## ############################################################

import re

## ############################################################

#   The following pattern defines variables as "${var-name}"
# where the variable name can be made of [A-Za-z0-9_-],
# i.e. alpha numerics plus "_" or "-"
#
VarPat = re.compile(r'\$\{([\w-]+)\}')

if __name__ == '__main__':

    d = {'var-1': 'value-1', 'var-2': 'value-2'}
    vstring = "A string with ${var-1} and ${var-2} and ${var-no-entry}"
    nstring = vstring

    for match in VarPat.finditer(vstring):
        mstr = match.group(1)
        if mstr in d:
            old = "${%s}" % (mstr)
            new = d[mstr]
            nstring = nstring.replace(old, new)

    print vstring
    print " (becomes)"
    print nstring
    

