package main

import "fmt"

func main() {
	data := []int{1, 2, 4, 7, 19}
	inOrder := InAscOrder(data)
	fmt.Println(inOrder)
}

func InAscOrder(numbers []int) bool {
	minNumber := numbers[0]
	for _, value := range numbers {
		if value < minNumber {
			return false
		}
		minNumber = value

	}
	return true
}
