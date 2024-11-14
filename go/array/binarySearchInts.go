// Дан отсортированный массив чисел и число N. Необходимо написать программу, которая использует встроенный алгоритм
// бинарного поиска, чтобы определить, присутствует ли N в данном массиве.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values, lookingValue := readInputData()
	isExists := searchValue(values, lookingValue)
	if isExists {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func searchValue(values []int, lookingValue int) bool {
	leftIdx, rightIdx := 0, len(values)-1
	for leftIdx <= rightIdx {
		middleIdx := (rightIdx + leftIdx) / 2
		middleValue := values[middleIdx]
		switch {
		case middleValue == lookingValue:
			return true
		case middleValue < lookingValue:
			leftIdx = middleIdx + 1
		case middleValue > lookingValue:
			rightIdx = middleIdx - 1
		}
	}
	return false
}

func readInputData() ([]int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for index := range size {
		scanner.Scan()
		values[index], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	lookingValue, _ := strconv.Atoi(scanner.Text())
	return values, lookingValue
}
