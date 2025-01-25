package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	syncStats := sync.Map{}
	var wg sync.WaitGroup

	for _, word := range words {

		wg.Add(1)
		go func() {
			defer wg.Done()
			amount := countDigits(word)
			syncStats.Store(word, amount)
		}()
	}

	wg.Wait()

	return asStats(&syncStats)
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

func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)
}
