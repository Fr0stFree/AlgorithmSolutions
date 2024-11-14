// У вас есть вектор целых чисел, который отсортирован в порядке возрастания. Напишите программу, которая принимает
// целое число N и использует встроенный алгоритм std::lower_bound для определения первого элемента в векторе,
// который не меньше N
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values, lookingValue := readInputData()
	idx := findLowerBound(values, lookingValue)
	fmt.Println(idx)
}

func findLowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low == 0 && array[low] > target {
		return -1
	}
	return low
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
