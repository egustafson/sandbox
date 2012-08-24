#!/bin/sh

# How to capture stderr in a back-tick command.

MYVAR=`(./printer | cat) 2> stderr.output`

echo "Beginning of shell script output"
echo "MYVAR = $MYVAR"