#!/bin/sh

DEFAULT_BOGUS_VAR=/bogus/dir

if [ "$BOGUS_VAR" = "" ]
then
	BOGUS_VAR=$DEFAULT_BOGUS_VAR
	export BOGUS_VAR
fi

echo "BOGUS_VAR = $BOGUS_VAR"
