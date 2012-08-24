#!/bin/sh


a_func () {
    if [ "$1" == "true" ]
    then
        true
    else
        false
    fi

    return $?
}


a_func "true"
echo $?
