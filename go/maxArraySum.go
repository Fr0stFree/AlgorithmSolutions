package main

import (
	"fmt"
)

func main() {
	example := []int{-3, 1, -8, 12, 2, -3, 5, -9, 4}
	result := findMaxSum(example, 2, 6)
	fmt.Println(result) // [12, 2]
}

func findMaxSum(numbers []int, L, R int) []int {
	var currentL, currentR, bestL, bestR, currentSum, bestSum int

	for index, value := range numbers[L:R] {
		if value > currentSum+value {
			currentSum = value
			currentL, currentR = index+L, index+L
		} else {
			currentSum += value
			currentR++
		}
		if currentSum > bestSum {
			bestSum = currentSum
			bestL, bestR = currentL, currentR
		}
	}
	return numbers[bestL : bestR+1]
}
