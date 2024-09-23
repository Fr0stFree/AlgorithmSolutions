package main

import (
	"fmt"
	"slices"
)

func main() {
	result := Solution("abcdefge")
	expected := []string{"ab", "cd", "ef"}
	fmt.Println(slices.Equal(result, expected))
}

func Solution(str string) []string {
	runeArray := []rune(str)
	var result []string

	for index := 0; index < len(runeArray); index += 2 {
		var nextLetter string
		if index+1 == len(runeArray) {
			nextLetter = "_"
		} else {
			nextLetter = string(runeArray[index+1])
		}
		result = append(result, string(runeArray[index])+nextLetter)
	}
	return result
}
