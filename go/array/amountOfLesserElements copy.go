// Создайте программу, которая выполняет бинарный поиск для определения количества
// элементов в отсортированном векторе, которые не превышают заданного значения.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	array, values := readInputData()
	for _, value := range values {
		amount := findAmountOfLesserValues(array, value)
		fmt.Println(amount)
	}
}

func findAmountOfLesserValues(values []int, threshold int) int {
	leftIdx, rightIdx := 0, len(values)-1
	var middleIdx int
	for leftIdx <= rightIdx {
		middleIdx = (rightIdx + leftIdx) / 2
		middleValue := values[middleIdx]
		if middleValue <= threshold {
			leftIdx = middleIdx + 1
		} else {
			rightIdx = middleIdx - 1
		}
	}
	amount := len(values[:middleIdx])
	if values[middleIdx] <= threshold {
		amount++
	}
	return amount
}

func readInputData() ([]int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	array := make([]int, size)
	for index := range size {
		scanner.Scan()
		array[index], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for index := range size {
		scanner.Scan()
		values[index], _ = strconv.Atoi(scanner.Text())
	}
	return array, values
}
