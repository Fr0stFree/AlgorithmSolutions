package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	height, width := readInputData()
	matrix := makeDiagonalMatrix(height, width)
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%3d", value)
		}
		fmt.Printf("\n")
	}
}

func makeDiagonalMatrix(height, width int) [][]int {
	var counter int
	matrix := make([][]int, height)
	for rowIdx := range matrix {
		matrix[rowIdx] = make([]int, width)
	}

	for diagIdx := 0; diagIdx < width+height; diagIdx++ {
		for rowIdx := 0; rowIdx <= diagIdx; rowIdx++ {
			if rowIdx >= height {
				continue
			}
			columnIdx := diagIdx - rowIdx
			if columnIdx >= width {
				continue
			}
			matrix[rowIdx][columnIdx] = counter
			counter++
		}
	}
	return matrix
}

func readInputData() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ")
	height, _ := strconv.Atoi(numbers[0])
	width, _ := strconv.Atoi(numbers[1])
	return height, width
}
