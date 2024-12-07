// Кролику нравится суммировать числа, поэтому он сложит числа во всех ячейках, куда будет прыгать.
// Какую максимальную сумму может получить кролик, когда из ячейки 0 допрыгает в ячейку n?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tests := readInputData()
	for _, test := range tests {
		fmt.Println(calcMaxWeight(test.gold, test.backpackSize))
	}
}

func calcMaxWeight(gold []int, maxSize int) int {
	var calc func(int, int) int
	memo := make(map[[2]int]int)

	calc = func(n, sizeLeft int) int {
		if n == 0 || sizeLeft == 0 {
			return 0
		}
		var result int
		if sizeLeft < gold[n-1] {
			result = calc(n-1, sizeLeft)
		} else {
			take := calc(n-1, sizeLeft-gold[n-1]) + gold[n-1]
			skip := calc(n-1, sizeLeft)
			result = max(take, skip)
		}
		memo[[2]int{n, sizeLeft}] = result
		return result
	}
	return calc(len(gold), maxSize)
}

type TestCase struct {
	gold 	   []int
	backpackSize int
}

func readInputData() []TestCase {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	testsAmount, _ := strconv.Atoi(scanner.Text())
	tests := make([]TestCase, testsAmount)
	for testIdx := range testsAmount {
		scanner.Scan()
		goldAmount, _ := strconv.Atoi(scanner.Text())
		gold := make([]int, goldAmount)

		scanner.Scan()
		backpackSize, _ := strconv.Atoi(scanner.Text())

		for idx := range goldAmount {
			scanner.Scan()
			gold[idx], _ = strconv.Atoi(scanner.Text())
		}
		tests[testIdx] = TestCase{gold: gold, backpackSize: backpackSize}
	}
	return tests
}
