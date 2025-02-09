package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

type Worker struct {
	fn         func() error
	ctx        context.Context
	cancel     context.CancelCauseFunc
	hasStarted bool
}

func NewWorker(fn func() error) *Worker {
	ctx, cancel := context.WithCancelCause(context.Background())
	return &Worker{fn: fn, ctx: ctx, cancel: cancel, hasStarted: false}
}

func (w *Worker) Start() {
	if w.hasStarted {
		return
	}
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
				err := w.fn()
				if err != nil {
					w.cancel(ErrFailed)
					return
				}
			}
		}
	}()
	w.hasStarted = true
}

func (w *Worker) Stop() {
	w.cancel(ErrManual)
}

func (w *Worker) AfterStop(fn func()) {
	if w.hasStarted {
		return
	}
	context.AfterFunc(w.ctx, fn)
}

func (w *Worker) Err() error {
	return context.Cause(w.ctx)
}

func main() {
	{
		count := 9
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(105 * time.Millisecond)
		worker.Stop()

		fmt.Println()
		// 9 8 7 6 5 4 3 2 1 0
	}
	{
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			if count == 0 {
				return errors.New("count is zero")
			}
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(35 * time.Millisecond)
		worker.Stop()

		fmt.Println(worker.Err())
		// 3 2 1 failed
	}
	{
		fn := func() error { return nil }

		worker := NewWorker(fn)
		worker.AfterStop(func() {
			fmt.Println("called after stop")
		})

		worker.Start()
		worker.Stop()

		time.Sleep(10 * time.Millisecond)
		// called after stop
	}
}
