// Дан массив, состоящий из целых чисел. Напишите программу, которая в данном массиве определит количество элементов,
// у которых два соседних и, при этом, оба соседних элемента меньше данного.
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
	result := countMaxima(data)
	fmt.Println(result)
}

func countMaxima(data []int) int {
	var counter int
	for index := 1; index < len(data)-1; index++ {
		if data[index] > data[index+1] && data[index] > data[index-1] {
			counter++
			index++
		}
	}
	return counter
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
