package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(676 == FindNextSquare(625))
	fmt.Println(-1 == FindNextSquare(114))
}

func FindNextSquare(sq int64) int64 {
	root := math.Sqrt(float64(sq))
	if float64(int(root)) != root {
		return -1
	}
	return int64(math.Pow(root+1, 2))
}
