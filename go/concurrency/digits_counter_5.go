package main

import (
	"fmt"
)

func count(cancel <-chan struct{}, start int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; ; i++ {
			select {
			case out <- i:
				continue
			case <-cancel:
				close(out)
				return
			}

		}
	}()
	return out
}

func take(cancel <-chan struct{}, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			select {
			case <-cancel:
				break
			case out <- <-in:
				continue
			}
		}
		close(out)
	}()
	return out
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	stream := take(cancel, count(cancel, 10), 5)
	first := <-stream
	second := <-stream
	third := <-stream

	fmt.Println(first, second, third)
}
