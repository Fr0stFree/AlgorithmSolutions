// Зайчик прыгает по массиву. Изначально он стоит в ячейке 0, а хочет попасть в ячейкуn.
// За один прыжок зайчик может прыгнуть на 1, на 3 или на 4 ячейки вправо.
// Прыгать он умеет только вправо.Сколько способов у него припрыгать в ячейку n?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	height, width := readInputData()
	amount := calcPathsAmount(height, width)
	fmt.Println(amount)
}

func calcPathsAmount(height, width int) int {
	var calc func(height, width int) int
	var memo = map[string]int{
		"1-1": 1,
	}

	calc = func(height, width int) int {
		if height <= 0 || width <= 0 {
			return 0
		}
		key := fmt.Sprintf("%d-%d", height, width)
		result, isCalculated := memo[key]
		if isCalculated {
			return result
		}
		memo[key] = calc(height-1, width) + calc(height, width-1)
		return memo[key]
	}
	return calc(height, width)
}

func readInputData() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	height, _ := strconv.Atoi(scanner.Text())
	return height, width
}
