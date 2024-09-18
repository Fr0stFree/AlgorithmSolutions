package main

import (
	"fmt"
)

func main() {
	myArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	result := CountPositivesSumNegatives(myArray[:])
	fmt.Println(result)
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	positiveCounter, negativeSum := 0, 0
	for _, value := range numbers {
		if value > 0 {
			positiveCounter++
		} else {
			negativeSum -= value
		}
	}
	return append(res, positiveCounter, -negativeSum)
}
