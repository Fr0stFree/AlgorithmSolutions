package main

import (
	"fmt"
	"strings"
)

func main() {
	res, _ := SizeToNumber("xxxl")
	fmt.Println(res)
}

func SizeToNumber(size string) (int, bool) {
	var sizes = map[string]int{
		"s": 36,
		"m": 38,
		"l": 40,
	}
	intSize, exists := sizes[size]
	if exists {
		return intSize, true
	}
	multiplier, specifiedSize := strings.Count(size, "x"), strings.TrimLeft(size, "x")
	intSize, _ = sizes[specifiedSize]
	if specifiedSize == "s" {
		return intSize - multiplier*2, true
	}
	if specifiedSize == "l" {
		return intSize + multiplier*2, true
	}
	return 0, false
}
