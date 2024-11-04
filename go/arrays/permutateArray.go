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
	result := makePermutedArray(data)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

func makePermutedArray(data []int) []int {
	result := make([]int, len(data))
	for index, value := range data {
		result[value-1] = index + 1
	}
	return result
}

func readInputData() []int {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	text, _ := reader.ReadString('\n')
	numbers := strings.Split(strings.TrimRight(text, "\n"), " ")

	result := make([]int, len(numbers))
	for index, elem := range numbers {
		result[index], _ = strconv.Atoi(elem)
	}
	return result
}
