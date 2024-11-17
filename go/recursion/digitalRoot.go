package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DigitalRoot(942) == 6)
	fmt.Println(DigitalRoot(132189) == 6)

}

func DigitalRoot(n int) int {
	digits := strings.Split(strconv.FormatInt(int64(n), 10), "")
	sum := 0
	for _, digit := range digits {
		number, _ := strconv.Atoi(digit)
		sum += number
	}å
	if sum > 9 {
		return DigitalRoot(sum)
	}
	return sum
}
