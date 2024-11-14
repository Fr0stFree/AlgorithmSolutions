package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const precision = 3
	value := readInputData()
	result := approximateSquareRoot(float64(value), precision)
	fmt.Println(strconv.FormatFloat(result, 'f', precision, 64))

}

func approximateSquareRoot(value float64, precision int) float64 {
	var right float64 = value
	var left float64 = 1
	var middle = (left + right) / 2

	for right-left > math.Pow(10, float64(-precision)) {
		middle = (left + right) / 2
		poweredValue := middle * middle
		switch {
		case poweredValue > value:
			right = middle
		case poweredValue < value:
			left = middle
		default:
			return middle
		}
	}
	return middle
}

func readInputData() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	return value
}
