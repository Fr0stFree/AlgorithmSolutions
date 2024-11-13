// Дан двумерный массив массив N×M. Требуется повернуть его по часовой стрелке на 90 градусов.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := readInputData()
	rotatedMatrix := rotate(matrix)
	fmt.Printf("%d %d\n", len(rotatedMatrix), len(rotatedMatrix[0]))
	for _, row := range rotatedMatrix {
		fmt.Printf("%s\n", strings.Trim(fmt.Sprint(row), "[]"))
	}
}

func rotate(matrix [][]int) [][]int {
	rotatedMatrix := make([][]int, len(matrix[0]))
	for rowIdx := range rotatedMatrix {
		rotatedMatrix[rowIdx] = make([]int, len(matrix))
	}

	for rowIdx, row := range matrix {
		for columnIdx, value := range row {
			rotatedMatrix[columnIdx][len(matrix)-rowIdx-1] = value
		}
	}
	return rotatedMatrix
}

func readInputData() [][]int {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	maybeSize := strings.Split(strings.TrimRight(text, "\n"), " ")
	rowsAmount, _ := strconv.Atoi(maybeSize[0])

	matrix := make([][]int, rowsAmount)
	for rowIdx := 0; rowIdx < rowsAmount; rowIdx++ {
		text, _ = reader.ReadString('\n')
		maybeNumbers := strings.Split(strings.TrimRight(text, "\n"), " ")
		row := make([]int, len(maybeNumbers))
		for idx, value := range maybeNumbers {
			row[idx], _ = strconv.Atoi(value)
		}
		matrix[rowIdx] = row
	}

	return matrix
}
