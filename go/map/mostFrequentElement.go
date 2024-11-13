package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"math"
)

func main() {
	data := readInput()
	frequency := findMostFrequentValue(data)
	fmt.Println(frequency)
}

func findMostFrequentValue[T int | string](values []T) T {
	frequencyMap := make(map[T]int)
	var mostFrequent T
	bestFrequency := math.MinInt

	for _, value := range values {
		frequencyMap[value]++
	}
	for key, frequency := range frequencyMap {
		if frequency > bestFrequency {
			bestFrequency, mostFrequent = frequency, key
		}
	}
	return mostFrequent
}

func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for idx := range size {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return values
}