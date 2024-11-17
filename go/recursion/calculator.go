package main

import (
	"fmt"
	"strings"
)

func main() {
	var target int
	fmt.Scan(&target)
	numbers := calculate(target)
	fmt.Println(len(numbers)-1)
	fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))
}

func calculate(target int) []int {
	var calculateInner func(target int) []int
	var memo = map[int][]int{
		1: make([]int, 0),
	}

	calculateInner = func(target int) []int {
		result, hasMemorized := memo[target]
		if hasMemorized {
			return result
		}
		if target < 1 {
			return nil
		}
		var options [][]int
		if target % 3 == 0 {
			option := append(make([]int, 0), calculateInner(target/3)...)
			option = append(option, target/3)
			options = append(options, option)
		}
		if target % 2 == 0 {
			option := append(make([]int, 0), calculateInner(target/2)...)
			option = append(option, target/2)
			options = append(options, option)
		}
		option := append(make([]int, 0), calculateInner(target-1)...)
		option = append(option, target-1)
		options = append(options, option)
		
		smallestResult := options[0]
		for _, option = range options {
			if len(option) < len(smallestResult) {
				smallestResult = option
			}
		}
		memo[target] = smallestResult
		return smallestResult
	}
	return append(calculateInner(target), target)
}