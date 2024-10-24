#!/bin/bash
function solve() {
    echo "$@" | tr -d "aAeEiIoOuU" && exit
}

solve "$@"