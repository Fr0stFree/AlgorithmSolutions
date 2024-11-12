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
	bubbleSort(values)
	fmt.Println(strings.Trim(fmt.Sprint(values), "[]"))
}

func bubbleSort(values []int) {
	for idx := range values {
		for innerIdx := 0; innerIdx < len(values) - idx - 1; innerIdx++ {
			if values[innerIdx] > values[innerIdx+1] {
				values[innerIdx], values[innerIdx+1] = values[innerIdx+1], values[innerIdx]
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
