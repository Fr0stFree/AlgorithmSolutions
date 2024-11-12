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
	insertionSort(values)
	fmt.Println(strings.Trim(fmt.Sprint(values), "[]"))
}

func insertionSort(values []int) {
	for idx := 1; idx < len(values); idx++ {
		targetValue, targetIdx := values[idx], idx - 1
		for targetIdx >= 0 && values[targetIdx] > targetValue {
			values[targetIdx+1] = values[targetIdx]
			targetIdx--
		}
		values[targetIdx+1] = targetValue
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
