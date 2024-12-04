// Дано натуральное число n, представьте его в виде суммы минимального количества кубов натуральных чисел.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	values := readInputData()
	results := decomposeCubes(values)
	for _, result := range results {
		fmt.Println(result)
	}
}

func decomposeCubes(values []int) []int {
	var do func (val float64) float64
	results := make([]int, len(values))
	memo := make(map[float64]float64)

	do = func (val float64) float64 {
		if val == 0 {
			return 0
		}
		possibleResult, isExists := memo[val]
		if isExists {
			return possibleResult
		}
		result := math.Inf(1)
		for counter := 1.0; math.Pow(counter, 3) <= val; counter++ {
			newResult := 1 + do(val - math.Pow(counter, 3))
			result = min(result, newResult)
		}
		memo[val] = result
		return result
	}

	for idx, value := range values {
		results[idx] = int(do(float64(value)))
	}
	return results
}

func readInputData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for idx := range size {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return values
}
