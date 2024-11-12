package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := readInput()
	result := mergeSort(values)
	fmt.Println(result)
}


func mergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	left, right := values[:len(values)/2], values[len(values)/2:]
	sortedLeft := mergeSort(left)
	sortedRight := mergeSort(right)
	result := merge(sortedLeft, sortedRight)
	return result
}

func merge(left, right []int) []int {
	var result []int
	var leftIdx, rightIdx int
	for {
		if left[leftIdx] < right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx++
			if leftIdx == len(left) {
				result = append(result, right[rightIdx:]...)
				break
			}
		} else {
			result = append(result, right[rightIdx])
			rightIdx++
			if rightIdx == len(right) {
				result = append(result, left[leftIdx:]...)
				break
			}
		}
	}
	return result
}

func readInput() []int {
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
