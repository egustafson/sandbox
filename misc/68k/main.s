#NO_APP
gcc2_compiled.:
___gnu_compiled_c:
.text
	.even
.globl _main
_main:
	jbsr ___main
	jbsr _get_sr
	addw #-1792,d0
	movew d0,a0
	movel a0,sp@-
	jbsr _set_sr
	moveq #0,d0
	addql #4,sp
	rts
	nop
	.even
.globl _set_sr
_set_sr:
#APP
	move sp@(6),sr
#NO_APP
	rts
	nop
	.even
.globl _get_sr
_get_sr:
#APP
	clrl d0
	move sr,d0
#NO_APP
	rts
	nop
