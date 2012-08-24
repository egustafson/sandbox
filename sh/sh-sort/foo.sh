#!/bin/sh
#
# The goal with this script is to create a merged, sorted, 
# unique list of files in multiple directories that start
# with "S??" (i.e. start files for init.d).
#

D1=d1
D2=d2

INITDIRS="$D1 $D2"

i=0
for d in $INITDIRS
do
  SCRIPTS=`find $d -name "S??*"`
  for s in $SCRIPTS 
  do
    f=`basename $s`
    INITF[$((i++))]=$f
  done
done

UNIQS=`echo ${INITF[*]} | tr ' ' '\n' | sort -u`

echo "-----"

echo "$UNIQS"

