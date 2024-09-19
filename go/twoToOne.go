package main

import (
	"fmt"
	"sort"
)

func main() {
	b := "xxxxyyyyabklmopq"
	a := "xyaabbbccccdefww"
	expected := "abcdefklmopqwxy"
	actual := TwoToOne(a, b)
	fmt.Println(expected == actual)
}

func Contains(letters []rune, letter rune) bool {
	for _, l := range letters {
		if letter == l {
			return true
		}
	}
	return false
}

func TwoToOne(s1 string, s2 string) string {
	var letters []rune
	for _, letter := range append([]rune(s1), []rune(s2)...) {
		if !Contains(letters, letter) {
			letters = append(letters, letter)
		}
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}
