// Создайте рекурсивную функцию countZeros, которая принимает целое число и возвращает количество нулей в его десятичной записи.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput()
	result := countZeros(data)
	fmt.Println(result)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text, _ = strings.CutSuffix(text, "\n")
	return text
}

func countZeros(number string) int {
	if len(number) == 0 {
		return 0
	}
	digit, _ := strconv.Atoi(string(number[0]))
	if digit == 0 {
		return 1 + countZeros(number[1:])
	}
	return countZeros(number[1:])
}
