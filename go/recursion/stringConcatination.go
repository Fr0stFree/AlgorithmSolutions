package main

import "fmt"

func main() {
	input := []string{"hello", "my", "dear", "friend"}
	result := concat(input, " ")
	fmt.Println(result)
}

func concat(strs []string, divider string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	return strs[0] + divider + concat(strs[1:], divider)
}
