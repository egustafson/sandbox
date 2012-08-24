#!/bin/ash

hickup () {
    echo "hickup..."
}

trap 'echo HUP' 1
trap 'echo INTR' 2
trap 'echo QUIT' 3

echo my pid $$

rm -f ./exit_file
while true
do
  if [ -e ./exit_file ]; then
      exit 0;
  fi
  echo "Starting child"
  ./child.ash
  echo "Child exited"
done
