#!/bin/sh

echo "Script start"

echo "ALL = $ALL"

for i in $ALL; do \
    echo file$i-*
done

echo "Script end"