// Даны числа n и m. Создайте массив A[n][m] и заполните его змейкой.
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	height, width := readInput()
	matrix := makeSnakeMatrix(height, width)
	var builder strings.Builder

	for idx, row := range matrix {
		for _, number := range row {
			builder.WriteString(fmt.Sprintf("%3d", number))
		}
		if idx != len(matrix)-1 {
			builder.WriteString("\n")
		}
	}
	fmt.Print(builder.String())
}

func makeSnakeMatrix(height, width int) [][]int {
	matrix := make([][]int, height)
	for rowIdx := range matrix {
		matrix[rowIdx] = make([]int, width)
	}
	for idx := 0; idx < height*width; idx++ {
		rowIdx, columnIdx := idx/width, idx%width
		matrix[rowIdx][columnIdx] = idx
		if columnIdx == width-1 && rowIdx%2 == 1 {
			slices.Reverse(matrix[rowIdx])
		}
	}
	return matrix
}
func readInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ")
	height, _ := strconv.Atoi(numbers[0])
	width, _ := strconv.Atoi(numbers[1])
	return height, width
}
