#!/bin/sh

if [ -f .kernel_dir ]; then
	if [ -d `cat .kernel_dir` ]; then
		exit 0;
	fi
	echo "'`cat .kernel_dir` does not exist, or is not a directory."
fi

if [ "$1" == "" ]; then
	echo ""
	echo " ---> You must set KERNEL_DIR=/path/to/kernel/build/dir the first time you build. <-- "
	echo ""
	exit 1;
else
	mkdir -p "$1"
	if [ -d $1 ]; then
		echo "KERNEL_DIR=$1" > .kernel_dir
	else
		echo "'$1' is not a valid directory"
		exit 1;
	fi
fi
