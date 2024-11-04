// Дан массив целых чисел. Преобразовать исходный массив, переставляя в обратном порядке элементы между максимальным и
// минимальным значениями массива, включая их.
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	index, data := readInputData()
	left, right := rotate(data[:len(data)-index]), rotate(data[len(data)-index:])
	result := rotate(slices.Concat(left, right))
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

func rotate(data []int) []int {
	result := make([]int, len(data))
	for index := range data {
		result[len(result)-1-index] = data[index]
	}
	return result
}

func readInputData() (int, []int) {
	reader := bufio.NewReader(os.Stdin)

	_, _ = reader.ReadString('\n')

	text, _ := reader.ReadString('\n')
	numbers := strings.Split(strings.TrimRight(text, "\n"), " ")

	result := make([]int, len(numbers))
	for index, elem := range numbers {
		result[index], _ = strconv.Atoi(elem)
	}

	text, _ = reader.ReadString('\n')
	shift, _ := strconv.Atoi(strings.TrimRight(text, "\n"))

	return shift, result
}
