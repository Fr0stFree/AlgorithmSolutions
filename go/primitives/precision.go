package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := readInputData()
	result := calculateSum(data)
	fmt.Println(strconv.FormatFloat(result, 'f', 20, 64))
}

func calculateSum(numbers []int) float64 {
	var straightSum int
	var inversedSum float64
	for _, value := range numbers {
		straightSum += value
		inversedSum += 1/float64(value)
	}
	return float64(straightSum) + inversedSum
}

func readInputData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	numbers := make([]int, size)
	for idx := range size {
		scanner.Scan()
		numbers[idx], _ = strconv.Atoi(scanner.Text())
	}
	return numbers
}
