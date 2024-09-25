#!/bin/bash
function longest() {
  echo "$1$2" | grep -o . | sort | uniq | tr -d "\n"
}

a="xyaabbbccccdefww"
b="xxxxyyyyabklmopq"
longest $a $b
