package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := readInput()
	selectionSort(values)
	fmt.Println(strings.Trim(fmt.Sprint(values), "[]"))
}

func selectionSort(values []int) {
	for idx := range values {
		for offset := range values[idx:] {
			if values[idx] > values[idx+offset] {
				values[idx], values[idx+offset] = values[idx+offset], values[idx]
			}
		}
	}
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
