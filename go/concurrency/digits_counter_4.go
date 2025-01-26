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


func submitWords(next nextFunc, out chan string) {
	defer close(out)
	for word := next(); word != ""; word = next() {
		out <- word
	}
}

func countWords(in chan string, out chan pair) {
	defer close(out)
	for word := range in {
		amount := countDigits(word)
		out <- pair{word, amount}
	}
}

func fillStats(in chan pair) counter {
	stats := make(map[string]int)
	for p := range in {
		stats[p.word] = p.count
	}
	return stats
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
