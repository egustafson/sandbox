#!/bin/sh

myfunc() {

    echo "Num Args: $#"
    echo "Arg1: $1"
    echo "Arg2: $2"
    echo "All args: $@"
}

echo
myfunc "xyzzy"
echo
myfunc "asdfg" 32
echo
myfunc 1 2 3 4 5 6
echo
