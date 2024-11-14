package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values, intervals := readInputData()
	for _, interval := range intervals {
		leftBorder := findLowerBound(values, interval[0])
		rightBorder := findUpperBound(values, interval[1])
		fmt.Println(rightBorder - leftBorder)
	}
}

func findLowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func findUpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func readInputData() ([]int, [][2]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for index := range size {
		scanner.Scan()
		values[index], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	intervals := make([][2]int, size)
	for index := range size {
		scanner.Scan()
		left, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())
		intervals[index] = [2]int{left, right}
	}
	return values, intervals
}
