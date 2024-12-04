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
	var memo = map[int]int{}

	find = func(currentIdx int) int {
		if currentIdx == 0 {
			return 1
		}
		if result, isExists := memo[currentIdx]; isExists {
			return result
		}
		memo[currentIdx] = 1
		for idx := currentIdx - 1; idx >= 0 && currentIdx-idx < 10; idx-- {
			if values[idx] < values[currentIdx] {
				memo[currentIdx] = max(memo[currentIdx], find(idx)+1)
			}
		}
		return memo[currentIdx]
	}

	maxLen := 0
	for idx := range values {
		maxLen = max(maxLen, find(idx))
	}
	return maxLen
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
