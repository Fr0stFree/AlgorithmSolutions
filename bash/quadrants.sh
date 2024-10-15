#!/bin/bash

function solve() {
    X=$1
    Y=$2
    if [ $X -gt 0 ] && [ $Y -gt 0 ]; then
        echo 1
    elif [ $X -gt 0 ] && [ $Y -lt 0 ]; then
        echo 4
    elif [ $X -lt 0 ] && [ $Y -gt 0 ]; then
        echo 2
    else
        echo 3
    fi
}

solve -10 100