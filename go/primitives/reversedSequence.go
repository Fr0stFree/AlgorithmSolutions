package main

import "fmt"

func main() {
	data := 10
	result := ReverseSeq(data)
	fmt.Println(result)
}

func ReverseSeq(n int) []int {
	var result []int
	for {
		result = append(result, n)
		n--
		if n == 0 {
			break
		}
	}
	return result
}
