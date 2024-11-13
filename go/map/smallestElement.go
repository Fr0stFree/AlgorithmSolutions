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
	value, hasFound := findSmallestValue(data)
	if hasFound {
		fmt.Println(value)
	} else {
		fmt.Println("NO")
	}
}

func findSmallestValue(values []int) (int, bool) {
	frequencyMap := make(map[int]int)
	smallestElement := math.MaxInt
	hasFound := false

	for _, value := range values {
		frequencyMap[value]++
	}
	for key, frequency := range frequencyMap {
		if key < smallestElement && frequency > 1 {
			smallestElement = key
			hasFound = true
		}
	}
	return smallestElement, hasFound
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