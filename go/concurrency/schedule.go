package main

import (
	"fmt"
	"time"
)

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	done := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fn()
			case <-done:
				return
			}
		}
	}()

	return func() {
		select {
		case <-done:
			return
		default:
			close(done)
		}
	}
}

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	time.Sleep(260 * time.Millisecond)
}
