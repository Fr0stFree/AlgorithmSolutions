package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

type Queue struct {
	elements chan int
}

func (q Queue) Get(block bool) (int, error) {
	if !block {
		select {
		case val := <-q.elements:
			return val, nil
		default:
			return 0, ErrEmpty
		}
	}
	return <-q.elements, nil
}

func (q Queue) Put(val int, block bool) error {
	if !block {
		select {
		case q.elements <- val:
		default:
			return ErrFull
		}
	} else {
		q.elements <- val
	}
	return nil
}

func MakeQueue(n int) Queue {
	elements := make(chan int, n)
	return Queue{elements}
}

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
