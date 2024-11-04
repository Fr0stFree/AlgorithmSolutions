package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	result := MinMax(data)
	fmt.Println(result)
}

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}
