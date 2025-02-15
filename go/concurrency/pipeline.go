package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generate(cancel <-chan struct{}) <-chan string{
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel: return
			case out <- randomWord(5):
			}
		}
	}()
	return out
}

func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string  {
	out := make(chan string)
	isUnique := func(word string) bool {
		letters := make(map[rune] bool)
		for _, letter := range word {
			isExist := letters[letter]
			if isExist {
				return false
			}
			letters[letter] = true
		}
		return true
	}
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel: return
			case word, ok := <-in:
				if !ok { return }
				unique := isUnique(word)
				if !unique { continue }
				select {
				case out <- word:
				case <-cancel: return
				}
			}
		}
	}()
	return out
}


func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	reverseWord := func(word string) string {
		letters := make([]rune, len(word))
		for idx, letter := range word {
			letters[len(word)-idx-1] = letter
		}
		reversedWord := string(letters)
		return fmt.Sprintf("%s -> %s\n", word, reversedWord)
	}
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel: return
			case word, ok := <- in:
				if !ok { return }
				reversed := reverseWord(word)
				select {
				case out <- reversed:
				case <-cancel: return
				}
			}
		}
	}()
	return out
}

func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	chans := []<-chan string{c1, c2}
	for _, in := range chans {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for {
				select {
				case <-cancel: return
				case out <- <-in:
				}
			}
		}()
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-cancel: return
		case word, ok := <- in:
			if !ok { return }
			fmt.Printf("%s", word)
		}
	}
}

func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
