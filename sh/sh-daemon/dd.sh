#!/bin/ash


echo "My pid $$" >/dev/tty

while true
do

  echo "stdout" 
  echo "stderr" 1>&2

  sleep 1

done

