# crt0.s          
#
# C RunTime - Necessary 'glue' code to link against a
#  gcc generated program.  
#
# Author:  Eric Gustafson
# Date:    24 February 1999
#

# ############################################################

# Initialize the initial stack pointer and program counter

.text
        .org    0x000000
_ISP:        
        .long   0x07ffe
_IPC:
        .long   _start

# Place the rest of the text segment starting at 0x0400

        .org    0x000400
.globl  _start
_start:
        bsr     _main
        .word   0x4848
# '0x4848' is a "break" instruction in the BSVC simulator.
# This implements a software break point. (i.e. 4E75 rts)

#
# Implement ___main for gcc.  Nothing to do, so just return.
#
.globl  ___main
___main:
        rts

# End
