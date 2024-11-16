package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) RemoveByValue(value int) {
	for idx, val := range *h {
		if val == value {
			heap.Remove(h, idx)
			break
		}
	}
}

func main() {
	carpool := &IntHeap{}
	heap.Init(carpool)
	guests := readInputData()
	order := resolveCarPooling(carpool, guests)
	for _, position := range order {
		fmt.Println(position)
	}
}

func resolveCarPooling(carpool *IntHeap, guests []Guest) []int {
	freePlaces := make(map[int]int, len(guests))
	for idx := 1; idx < len(guests); idx++ {
		freePlaces[idx] = -1
		heap.Push(carpool, idx)
	}
	var positions []int

	for _, guest := range guests {
		if guest.action == Incoming {
			isFree := freePlaces[guest.position] == -1
			if isFree {
				carpool.RemoveByValue(guest.position)
				freePlaces[guest.position] = guest.id
				positions = append(positions, guest.position)
			} else {
				position := heap.Pop(carpool).(int)
				freePlaces[position] = guest.id
				positions = append(positions, position)
			}
		} else {
			for position, id := range freePlaces {
				if id == guest.id {
					freePlaces[position] = -1
					heap.Push(carpool, position)
					break
				}
			}
		}

	}
	return positions
}

type Guest struct {
	id       int
	position int
	action   Action
}

type Action string

const (
	Leaving  Action = "-"
	Incoming Action = "+"
)

func readInputData() []Guest {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	guests := make([]Guest, size)
	for idx := range size {
		scanner.Scan()
		guestType := scanner.Text()
		scanner.Scan()
		position, _ := strconv.Atoi(scanner.Text())
		switch guestType {
		case "+":
			guests[idx] = Guest{idx + 1, position, Incoming}
		case "-":
			guests[idx] = Guest{position, -1, Leaving}
		default:
			panic("unknown guest type")
		}
	}

	return guests
}
