// Найдите самую длинную красивую подпоследовательность массива a и выведите её длину.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := readInputData()
	length := findLenOfLongestSubsequence(values)
	fmt.Println(length)
}

func findLenOfLongestSubsequence(values []int) int {
	var find func(int) int

	find = func(currentIdx int) int {
		if currentIdx == 1 {
			fmt.Printf("return 1 because currentIdx is 1\n")
			return 1
		}
		var options = []int{1}
		for i := currentIdx; i < 0 && currentIdx - i < 10; i-- {
			if values[i] > values[i-1] {
				option := 1 + find(i-1)
				options = append(options, option)
			}
			
		}
		fmt.Printf("options: %v\n", options)
		for _, option := range options {
			if option > options[0] {
				options[0] = option
			}
		}
		fmt.Printf("return options[0]: %d\n", options[0])
		return options[0]

	}
	return find(len(values)-1)
}

func readInputData() []int {
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
