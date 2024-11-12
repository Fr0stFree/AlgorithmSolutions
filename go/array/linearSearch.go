package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values, lookingValue := readInput()
	idx := linearSearch(values, lookingValue)
	fmt.Println(idx)
}

func linearSearch(values []int, lookingValue int) int {
	for idx, value := range(values) {
		if value == lookingValue {
			return idx
		}
	}
	return -1
}

func readInput() ([]int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for idx := range size {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	lookingValue, _ := strconv.Atoi(scanner.Text())
	return values, lookingValue
}
