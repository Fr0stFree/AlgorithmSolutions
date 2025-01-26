package main

import (
	"fmt"
	"strings"
)

func encode(str string) string {

	submitter := func(str string) <-chan string {
		ch := make(chan string)
		go func() {
			words := strings.Fields(str)
			for _, word := range words {
				ch <- word
			}
			close(ch)
		}()
		return ch
	}

	encoder := func(ch1 <-chan string) <-chan string {
		ch2 := make(chan string)
		go func() {
			for word := range ch1 {
				ch2 <- encodeWord(word)
			}
			close(ch2)
		}()
		return ch2
	}

	receiver := func(ch <-chan string) []string {
		words := []string{}
		for word := range ch {
			words = append(words, word)
		}
		return words
	}

	// конец решения

	pending := submitter(str)
	encoded := encoder(pending)
	words := receiver(encoded)
	return strings.Join(words, " ")
}

func encodeWord(word string) string {
	const shift = 13
	const char_a byte = 'a'
	encoded := make([]byte, len(word))
	for idx, char := range []byte(word) {
		delta := (char - char_a + shift) % 26
		encoded[idx] = char_a + delta
	}
	return string(encoded)
}

func main() {
	src := "go is awesome"
	res := encode(src)
	fmt.Println(res)
}
