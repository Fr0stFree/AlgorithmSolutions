package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := readInput()
	result := countingSort(values)
	fmt.Println(result)
}


func countingSort(values []int) []int {
	counts := make([]int, len(values))
	result := make([]int, 0)

	for _, value := range values {
		counts[value]++
	}
	for idx, value := range counts {
		for _ = range value {
			result = append(result, idx)
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
