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
	target := readInputData()
	amount := calcHopsAmount(target)
	fmt.Println(amount)
}

func calcHopsAmount(target int) int {
	var calculateHops func(target int) int
	var memo = map[int]int{
		1: 1,
		3: 2,
		4: 4,
	}

	calculateHops = func(target int) int {
		if target < 1 {
			return 0
		}
		result, isExist := memo[target]
		if isExist {
			return result
		}
		result = calculateHops(target-1) + calculateHops(target-3) + calculateHops(target-4)
		memo[target] = result
		return result
	}
	return calculateHops(target)
}

func readInputData() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	return target
}
