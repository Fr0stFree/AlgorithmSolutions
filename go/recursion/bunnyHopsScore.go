// Кролику нравится суммировать числа, поэтому он сложит числа во всех ячейках, куда будет прыгать.
// Какую максимальную сумму может получить кролик, когда из ячейки 0 допрыгает в ячейку n?
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	target := readInputData()
	amount := calcHopsScore(target)
	fmt.Println(amount)
}

func calcHopsScore(scores []int) int {
	var calculateScore func(target int) int
	var memo = map[int]int{0: scores[0]}

	calculateScore = func(target int) int {
		if target < 0 {
			return math.MinInt
		}
		score, isExist := memo[target]
		if isExist {
			return score
		}
		score = scores[target] + max(calculateScore(target-1), calculateScore(target-3), calculateScore(target-5))
		memo[target] = score
		return score
	}
	return calculateScore(len(scores)-1)
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
