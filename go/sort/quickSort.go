package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	values := readInput()
	quckSort(values)
	fmt.Println(strings.Trim(fmt.Sprint(values), "[]"))
}

func quckSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	pivot := values[len(values)/2]
	less, equal, bigger := partition(values, pivot)
	sortedLess := quckSort(less)
	sortedBigger := quckSort(bigger)
	return slices.Concat(sortedLess, equal, sortedBigger)
}

func partition(values []int, pivot int) ([]int, []int, []int) {
	var less, equal, bigger []int

	for _, value := range values {
		switch {
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			bigger = append(bigger, value)
		case value < pivot:
			less = append(less, value)
		}
	}
	return less, equal, bigger
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
