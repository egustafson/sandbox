#!/bin/sh

a="asdfg"
b="asdfg"

c="xyzzy"
d="xyzzy"

if [ $a = $b ]; then echo "a and b equal"; fi

if [ $c = $d ]; then echo "c and d equal"; fi

if [  $a = $b  -a $c = $d ]; then echo "all good"; fi