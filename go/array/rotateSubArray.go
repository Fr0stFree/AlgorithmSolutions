// Дан массив целых чисел. Преобразовать исходный массив, переставляя в обратном порядке элементы между максимальным и
// минимальным значениями массива, включая их.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := readInputData()
	_, curMaxIdx, _, curMinIdx := findMinAndMax(data)
	left, right := min(curMinIdx, curMaxIdx), max(curMinIdx, curMaxIdx)
	rotatedPart := rotate(data[left : right+1])
	result := slices.Concat(data[:left], rotatedPart, data[right+1:])
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

func findMinAndMax(data []int) (int, int, int, int) {
	var curMinIdx, curMaxIdx int
	curMax := math.MinInt
	curMin := math.MaxInt

	for index, value := range data {
		if value > curMax {
			curMax = value
			curMaxIdx = index
		}
		if value < curMin {
			curMin = value
			curMinIdx = index
		}
	}
	return curMax, curMaxIdx, curMin, curMinIdx
}

func rotate(data []int) []int {
	result := make([]int, len(data))
	for index := range data {
		result[len(result)-1-index] = data[index]
	}
	return result
}

func readInputData() []int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	numbers := strings.Split(strings.TrimRight(text, "\n"), " ")

	result := make([]int, len(numbers))
	for index, elem := range numbers {
		result[index], _ = strconv.Atoi(elem)
	}
	return result
}
