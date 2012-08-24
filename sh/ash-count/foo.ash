#!/bin/ash

for x in 1 3 5 7 9 
do
    ARR="$ARR $x"
done

echo "${ARR}"

##

FILES=`find d1 -name "S??*"`
for f in $FILES
do
    echo "file -> $f"
    INITF="$INITF `basename $f`"
done

CMDS=`echo $INITF | tr ' ' '\n' | sort -u`

for c in $CMDS
do
    echo "CMD = \"$c\""
done
