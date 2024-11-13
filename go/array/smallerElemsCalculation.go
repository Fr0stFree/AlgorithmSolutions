package main

import (
	"fmt"
	"slices"
)

func main() {
	actual := Smaller([]int{5, 4, 3, 2, 1})
	expected := []int{4, 3, 2, 1, 0}
	fmt.Println(slices.Compare(actual, expected))
}

func Smaller(arr []int) []int {
	var result = make([]int, len(arr))

	for index, value := range arr {
		var amount int
		for _, nextValue := range arr[index:] {
			if nextValue < value {
				amount++
			}
		}
		result[index] = amount

	}
	return result
}
