package main

import (
	"fmt"
	"sort"
	"strconv"
	"bufio"
	"os"
	"math"
)

func main() {
	data := readInput()
	minValue, maxValue := findMinMaxByLinearSearch(data)
	fmt.Println(minValue)
	fmt.Println(maxValue)
}

func findMinMaxBySort(arr []int) (int, int) {
	sort.Ints(arr)
	return arr[0], arr[len(arr)-1]
}

func findMinMaxByLinearSearch(arr []int) (int, int) {
	minValue, maxValue := math.MaxInt, math.MinInt
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
		if value < minValue {
			minValue = value
		}
	}
	return minValue, maxValue
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