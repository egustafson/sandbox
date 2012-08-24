#!/bin/sh

func1() {
    echo "func1"
    return 1
}

func2() {
    echo "func2"
    return 0
}

func3() {
    echo "func3"
    return -1
}

comb() {
    func1
    f1=$?
    echo "f1 = $f1"
    func2
    f2=$?
    echo "f2 = $f2"

    func3
    f3=$?
    echo "f3 = $f3"


    return $f1 || $f2
}

comb
ret=$?

echo "ret = $ret"
