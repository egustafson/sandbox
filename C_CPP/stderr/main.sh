#!/bin/sh

echo "printer > /dev/null"
echo "---------------------------"
printer > /dev/null
echo "---------------------------"
echo "printer > /dev/null 2>&1"
echo "---------------------------"
printer > /dev/null 2>&1
echo "---------------------------"
