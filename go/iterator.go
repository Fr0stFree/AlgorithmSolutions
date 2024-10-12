package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type element interface{}

type weightFunc func(element) int

type iterator interface {
	next() bool
	val() element
}

type intIterator struct {
	values  []int
	current int
}

func (i *intIterator) next() bool {
	if i.current >= len(i.values) {
		return false
	}
	return true
}

func (i *intIterator) val() element {
	val := i.values[i.current]
	i.current++
	return val
}

func newIntIterator(src []int) *intIterator {
	return &intIterator{src, 0}
}

func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
