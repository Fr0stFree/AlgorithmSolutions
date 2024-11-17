package main

import (
	"fmt"
)

func main() {
	sum := 9
	numbers := []int{2, 4, 7, 5}
	isPossible := isSumExist(sum, numbers)
	fmt.Println(isPossible)
}

func isSumExist(sum int, numbers []int) bool {
	var isExist func(sum int) (bool, []int)

	isExist = func(sum int) (bool, []int) {
		if sum == 0 {
			return true, make([]int, 0)
		}
		if sum < 0 {
			return false, nil
		}
		for _, number := range numbers {
			newSum := sum - number
			exists, summands := isExist(newSum)
			if exists {
				summands = append(summands, number)
				return true, summands
			}
		}
		return false, nil
	}
	exists, path := isExist(sum)
	fmt.Println(path)
	return exists
}