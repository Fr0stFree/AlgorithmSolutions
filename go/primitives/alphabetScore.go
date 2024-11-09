package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(High("man i need a taxi up to ubud") == "taxi")
	fmt.Println(High("aa b") == "aa")
}

func High(s string) string {
	result := struct {
		word  string
		score int32
	}{}

	for _, word := range strings.Split(s, " ") {
		score := calculateWordScore(word)
		if score > result.score {
			result.word = word
			result.score = score
		}
	}
	return result.word
}

func calculateWordScore(word string) int32 {
	var sum int32
	for _, letter := range word {
		sum += letter - 96
	}
	return sum
}
