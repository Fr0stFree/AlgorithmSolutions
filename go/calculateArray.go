// Дан массив. Преобразовать исходный массив, вычитая из значений отрицательных элементов массива количество положительных
// значений, а из значений положительных элементов массива количество отрицательных значений.
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
	elements := recalculate(data)
	fmt.Println(strings.Trim(fmt.Sprint(elements), "[]"))
}

func recalculate(data []int) []int {
	var negativeAmount, positiveAmount int

	for _, value := range data {
		if value > 0 {
			positiveAmount++
		} else if value < 0 {
			negativeAmount++
		}
	}

	for index, value := range data {
		if value > 0 {
			data[index] -= negativeAmount
		} else if value < 0 {
			data[index] -= positiveAmount
		}
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
