package main

import (
	"fmt"
	"time"
)

func gather(funcs []func() any) []any {
	results := make([]any, len(funcs))
	done := make(chan struct{})
	for idx, fn := range funcs {
		go func() {
			results[idx] = fn()
			done <- struct{}{}
		}()
	}
	for range funcs {
		<- done
	}
	return results
}

func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
