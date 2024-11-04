// Дан массив чисел. Преобразовать исходный массив, вычитая максимальное значение массива из элементов массива,
// идущих после минимального.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInputData()
	elements := solve(data)
	fmt.Println(strings.Trim(fmt.Sprint(elements), "[]"))
}

func solve(data []int) []int {
	var currentMin, currentMax, minIndex int

	for index, value := range data {
		if value > currentMax {
			currentMax = value
		}
		if value < currentMin {
			currentMin = value
			minIndex = index
		}
	}

	for index, _ := range data[minIndex+1:] {
		data[minIndex+index+1] -= currentMax
	}
	return data
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
