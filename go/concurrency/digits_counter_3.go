package main

import (
	"fmt"
	"strings"
	"unicode"
)

type nextFunc func() string

type counter map[string]int

type pair struct {
	word  string
	count int
}

func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

func submitWords(next nextFunc, pending chan string) {
	defer close(pending)
	for word := next(); word != ""; word = next() {
		pending <- word
	}
}

func countWords(pending chan string, counted chan pair) {
	defer close(counted)
	for word := range pending {
		amount := countDigits(word)
		counted <- pair{word, amount}
	}
}

func fillStats(counted chan pair) counter {
	results := make(map[string]int)
	for p := range counted {
		results[p.word] = p.count
	}
	return results
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
