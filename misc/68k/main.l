
===================  _start (crt0.o)

                        	.globl  _start
                        	_start:
00000400  6100 FBFE 	        bsr     _main
00000404  4848      	        .word   0x4848
                        	# '0x4848' is a "break" instruction in the BSVC simulator.
                        	# This implements a software break point. (i.e. 4E75 rts)

===================  ___main (crt0.o)

                        	.globl  ___main
                        	___main:
00000406  4E75      	        rts

===================  _main (main.o)

                        	.globl _main
                        	_main:
00000408  4EB9 0000 		jbsr ___main
               0000 
0000040e  611A      		jbsr _get_sr
00000410  0640 F900 		addw #-1792,d0
00000414  3040      		movew d0,a0
00000416  2F08      		movel a0,sp@-
00000418  6108      		jbsr _set_sr
0000041a  7000      		moveq #0,d0
0000041c  588F      		addql #4,sp
0000041e  4E75      		rts

===================  _set_sr (main.o)

                        	.globl _set_sr
                        	_set_sr:
                        	#APP
00000422  46EF 0006 		move sp@(6),sr
                        	#NO_APP
00000426  4E75      		rts

===================  _get_sr (main.o)

                        	.globl _get_sr
                        	_get_sr:
                        	#APP
0000042a  4280      		clrl d0
0000042c  40C0      		move sr,d0
                        	#NO_APP
0000042e  4E75      		rts
