package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IsTriangle(1, 2, 3) == false)
	fmt.Println(IsTriangle(-1, -1, -1) == true)
}

func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	p := 0.5 * float64(a+b+c)
	s := math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))
	fmt.Println(s)
	return s > 0
}
