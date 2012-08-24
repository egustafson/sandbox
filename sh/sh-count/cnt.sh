#!/bin/ash

ii=0

files=`find . -type f`

for f in $files 
do
	ii=$(($ii+1))
done

echo "There are $ii files in the directory"
