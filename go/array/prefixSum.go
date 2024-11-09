package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	array, indexes := readInputData()
	result := calculateSums(array, indexes)
	fmt.Println(result)
}

func calculateSums(array []int, indexes [][2]int) int {
	var calculatedSum int
	prefixArray := make([]int, len(array)+1)
	for index, value := range array {
		prefixArray[index+1] = prefixArray[index] + value
	}
	for _, index := range indexes {
		sum := prefixArray[index[1]] - prefixArray[index[0]-1]
		calculatedSum += sum
	}
	return calculatedSum
}

func readInputData() ([]int, [][2]int) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	metadata := strToIntArray(input)

	input, _ = reader.ReadString('\n')
	array := strToIntArray(input)

	results := make([][2]int, 0)
	for index := 0; index < metadata[1]; index++ {
		input, _ = reader.ReadString('\n')
		numbers := strToIntArray(input)
		results = append(results, [2]int{numbers[0], numbers[1]})
	}
	return array, results
}

func strToIntArray(data string) []int {
	strData, _ := strings.CutSuffix(data, "\n")
	strArray := strings.Split(strData, " ")
	intArray := make([]int, len(strArray))
	for index, value := range strArray {
		intArray[index], _ = strconv.Atoi(value)
	}
	return intArray
}
