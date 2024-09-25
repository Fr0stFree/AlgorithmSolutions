#!/bin/bash
length=$1
width=$2
height=$3
result=$(bc<<<"scale=3;$length*$width*$height")
echo $result
