package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "8 3 -5 42 -1 0 0 -9 4 7 4 -4"
	expected := "42 -9"
	result := HighAndLow(input)
	fmt.Println(result == expected)
}

func HighAndLow(in string) string {
	strNumbersArray := strings.Split(in, " ")
	result := make([]int, len(strNumbersArray))

	for index, strNumber := range strNumbersArray {
		number, _ := strconv.Atoi(strNumber)
		result[index] = number
	}
	sort.Ints(result)
	return fmt.Sprintf("%d %d", result[len(result)-1], result[0])
}
