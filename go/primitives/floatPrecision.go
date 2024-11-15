package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	result := calculateFuncDiff(readInputData())
	fmt.Println(strconv.FormatFloat(result, 'f', 20, 64))
}

func calculateFuncDiff(a, b, x1, x2 int) float64 {
	result := float64(a) + 1/float64(b)
	if x1 < x2 {
		return -result
	}
	return result
}

func readInputData() (int, int, int, int) {
	const size = 4
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	var result [size]int
	for idx := range size {
		scanner.Scan()
		result[idx], _ = strconv.Atoi(scanner.Text())
	}
	
	return result[0], result[1], result[2], result[3]
}
