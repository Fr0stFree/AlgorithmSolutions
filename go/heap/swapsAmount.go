package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type CustomHeap struct {
	container []int
}

func (h *CustomHeap) Elem(idx int) int {
	return h.container[idx]
}

func (h *CustomHeap) HasRoot(idx int) bool {
	return idx > 1
}

func (h *CustomHeap) Root(idx int) (int, int) {
	rootIdx := idx / 2
	return h.container[rootIdx], rootIdx
}

func (h *CustomHeap) HasLeft(idx int) bool {
	return h.Len() >= 2*idx
}

func (h *CustomHeap) Left(idx int) (int, int) {
	leftIdx := 2 * idx
	return h.container[leftIdx], leftIdx
}

func (h *CustomHeap) HasRight(idx int) bool {
	return h.Len() >= 2*idx+1
}

func (h *CustomHeap) Right(idx int) (int, int) {
	rightIdx := 2*idx + 1
	return h.container[rightIdx], rightIdx
}

func (h *CustomHeap) Swap(firstIdx, secondIdx int) {
	h.container[firstIdx], h.container[secondIdx] = h.container[secondIdx], h.container[firstIdx]
}

func (h *CustomHeap) Len() int {
	return len(h.container) - 1
}

func MakeCustomHeap(elements []int) *CustomHeap {
	return &CustomHeap{slices.Concat([]int{0}, elements)}
}

func main() {
	data := readInputData()
	heap := MakeCustomHeap(data)
	swaps := calcSwapAmount(*heap)

	var builder strings.Builder
	builder.WriteString(fmt.Sprintln(len(swaps)))
	for _, swap := range swaps {
		builder.WriteString(fmt.Sprintf("%d %d\n", swap[0], swap[1]))
	}
	fmt.Print(builder.String())
}

func calcSwapAmount(heap CustomHeap) [][2]int {
	idx := 1
	var swaps [][2]int

	for idx <= heap.Len() {
		value := heap.Elem(idx)
		if heap.HasLeft(idx) {
			leftValue, leftIdx := heap.Left(idx)
			fmt.Printf("left elem of %d is %d\n", idx, leftIdx)
			if leftValue < value {
				swaps = append(swaps, [2]int{leftIdx, idx})
				fmt.Printf("should swap %d and %d\n", leftIdx, idx)
				heap.Swap(leftIdx, idx)
				fmt.Printf("checking if %d has root\n", idx)
				if heap.HasRoot(idx) {
					_, rootIdx := heap.Root(idx)
					fmt.Printf("%d has root %d\n", idx, rootIdx)
					idx = rootIdx
					continue
				}
			}
		}
		if heap.HasRight(idx) {
			rightValue, rightIdx := heap.Right(idx)
			fmt.Printf("right elem of %d is %d\n", idx, rightIdx)
			if rightValue < value {
				swaps = append(swaps, [2]int{rightIdx, idx})
				fmt.Printf("should swap %d and %d\n", rightIdx, idx)
				heap.Swap(rightIdx, idx)
				fmt.Printf("checking if %d has root\n", idx)
				if heap.HasRoot(idx) {
					_, rootIdx := heap.Root(idx)
					fmt.Printf("%d has root %d\n", idx, rootIdx)
					idx = rootIdx
					continue
				}
			}
		}
		idx++
	}
	return swaps
}

func readInputData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for idx := range size {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return values
}
