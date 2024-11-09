package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputData := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	expectedResult := [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}
	fmt.Println(expectedResult)
	fmt.Println(solve(inputData))

}

func solve(words []string) [][]string {
	storage := make(map[string][]string, len(words))

	for _, word := range words {
		splitWord := strings.Split(word, "")
		sort.Strings(splitWord)
		key := strings.Join(splitWord, "")
		anagrams := storage[key]
		storage[key] = append(anagrams, word)
	}

	result := make([][]string, len(storage))
	index := 0
	for _, anagrams := range storage {
		result[index] = anagrams
		index++
	}

	return result
}
