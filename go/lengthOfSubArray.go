package main

import "fmt"

func main() {
	array := []int{0, -3, 7, 4, 0, 3, 7, 9}
	number := 7
	fmt.Println(LengthOfSequence(array, number) == 5)
}

func LengthOfSequence(arr []int, key int) int {
	start, finish := -1, -1

	for index, value := range arr {
		if value != key {
			continue
		}
		if start == -1 {
			start = index
			continue
		}
		if finish == -1 {
			finish = index
			continue
		}
		return 0
	}
	if start == -1 || finish == -1 {
		return 0
	}
	return len(arr[start : finish+1])
}
