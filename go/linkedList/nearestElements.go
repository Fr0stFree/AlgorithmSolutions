package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := readInputData()
	result := buildArrayOfNearest(values)
	for _, pair := range result {
		fmt.Printf("%d %d\n", pair[0], pair[1])
	}
}
func buildArrayOfNearest(values []int) [][2]int {
	nearest := make([][2]int, len(values))
	deque := list.New()

	for idx, value := range values {
		var toRemove []*list.Element

		for prevElem := deque.Back(); prevElem != nil; prevElem = prevElem.Prev() {
			if prevElem.Value.(int) >= value {
				toRemove = append(toRemove, prevElem)
				continue
			}
			nearest[idx][0] = prevElem.Value.(int)
			break
		}
		for _, elem := range toRemove {
			deque.Remove(elem)
		}
		deque.PushFront(value)
	}

	deque.Init()
	for idx := len(values) - 1; idx != -1; idx-- {
		var toRemove []*list.Element

		for nextElem := deque.Back(); nextElem != nil; nextElem = nextElem.Prev() {
			if nextElem.Value.(int) >= values[idx] {
				toRemove = append(toRemove, nextElem)
				continue
			}
			nearest[idx][1] = nextElem.Value.(int)
			break
		}
		for _, elem := range toRemove {
			deque.Remove(elem)
		}
		deque.PushBack(values[idx])
	}

	return nearest
}

func readInputData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())
	values := make([]int, amount)
	for idx := range amount {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return values
}
