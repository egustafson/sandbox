#!/bin/sh

byref () {

    echo "Setting variable $1 to have value $2"
    eval "$1=$2"
}

byref MYVAR 12345


echo "MYVAR = $MYVAR"