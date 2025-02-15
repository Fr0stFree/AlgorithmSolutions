package main

import (
	"context"
	"fmt"
	"strings"
	"unicode"
)

type counter map[string]int

type pair struct {
	word  string
	count int
}

func countDigitsInWords(ctx context.Context, words []string) counter {
	pending := submitWords(ctx, words)
	counted := countWords(ctx, pending)
	return fillStats(ctx, counted)
}

func submitWords(ctx context.Context, words []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, word := range words {
			select {
			case out <- word:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func countWords(ctx context.Context, in <-chan string) <-chan pair {
	out := make(chan pair)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case word, ok := <-in:
				if !ok {
					return
				}
				count := countDigits(word)
				select {
				case out <- pair{word, count}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func fillStats(ctx context.Context, in <-chan pair) counter {
	stats := counter{}
	for {
		select {
		case <-ctx.Done():
			return stats
		case p, ok := <-in:
			if !ok {
				return stats
			}
			stats[p.word] = p.count
		}
	}
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

func main() {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stats := countDigitsInWords(ctx, words)
	fmt.Println(stats)
}
