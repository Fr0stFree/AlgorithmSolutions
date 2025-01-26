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

	done := make(chan struct{})
	counted := make(chan pair)
	
	for i := 0; i < 4; i++ {
		go countWords(done, pending, counted)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-done
		}
		close(counted)
	}()

	return fillStats(counted)
}

func submitWords(next nextFunc, out chan<- string) {
	for {
		word := next()
		if word == "" {
			break
		}
		out <- word
	}
	close(out)
}

func countWords(done chan<- struct{}, in <-chan string, out chan<- pair) {
	for word := range in {
		amount := countDigits(word)
		out <- pair{word, amount}
	}
	done <- struct{}{}
}

func fillStats(in <-chan pair) counter {
	stats := counter{}
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
	phrase := "1 22 333 4444 55555 666666 7777777 88888888"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
