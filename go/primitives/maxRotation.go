package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input int64 = 56789
	const expected int64 = 68957
	var result = MaxRot(input)
	fmt.Println(result == expected)
}

func MaxRot(n int64) int64 {
	splitNumbers := strings.Split(strconv.Itoa(int(n)), "")
	result, _ := strconv.Atoi(strings.Join(splitNumbers, ""))

	for index := 0; index < len(splitNumbers); index++ {
		rotated, keep := append(splitNumbers[index+1:], splitNumbers[index]), splitNumbers[:index]
		number, _ := strconv.Atoi(strings.Join(append(keep, rotated...), ""))
		if number > result {
			result = number
		}
	}
	return int64(result)
}
