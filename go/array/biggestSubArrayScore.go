package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	values := readInput()
	result := findBiggestScore(values)
	fmt.Println(result)
}

func findBiggestScore(values []int) int {
	maxScore := math.MinInt
	for idx, value := range values {
		score := value
		for leftIdx := idx - 1; leftIdx > 0; leftIdx-- {
			if values[leftIdx] < value {
				break
			}
			score += value
		}
		for rightIdx := idx + 1; rightIdx < len(values); rightIdx++ {
			if values[rightIdx] < value {
				break
			}
			score += value
		}
		maxScore = max(maxScore, score)
	}
	return maxScore
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
