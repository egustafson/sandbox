#!/bin/sh

if [ $# -ne 1 ]; then
    echo ""
    echo "Usage:  $0 <version>"
    echo ""
    echo "   example:  $0 v9.8.7r06"
    echo ""
    exit
fi


echo "Version:  $1"

