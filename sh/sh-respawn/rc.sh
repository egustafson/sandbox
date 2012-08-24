#!/bin/sh

DAEMON=./daemon.sh

if [ -x wrapper.sh ]
then
    echo "Starting..."
    (while true; do echo "Starting daemon.sh"; $DAEMON; sleep 1; done)&
else
    echo "Starting...FAILED"
fi

echo "rc.sh done."
