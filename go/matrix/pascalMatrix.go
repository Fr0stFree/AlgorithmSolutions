// Даны два числа n и m. Создайте двумерный массив A[n][m] и заполните его по следующим правилам: Числа, стоящие в
// строке 0 или в столбце 0 равны 1 (A[0][j]=1, A[i][0]=1). Для всех остальных элементов массива
// A[i][j]=A[i-1][j]+A[i][j-1], то есть каждый элемент равен сумме двух элементов, стоящих слева и сверху от него.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	width, height := readInputData()
	matrix := makePascalMatrix(width, height)
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("     %d", value)
		}
		fmt.Printf("\n")
	}
}

func makePascalMatrix(width, height int) [][]int {
	matrix := make([][]int, height)
	for rowIdx := range matrix {
		matrix[rowIdx] = make([]int, width)
		matrix[rowIdx][0] = 1
	}
	for columnIdx := 0; columnIdx < width; columnIdx++ {
		matrix[0][columnIdx] = 1
	}
	for rowIdx, row := range matrix[1:] {
		for columnIdx := range row[1:] {
			amount := matrix[rowIdx][columnIdx+1] + matrix[rowIdx+1][columnIdx]
			matrix[rowIdx+1][columnIdx+1] += amount
		}
	}
	return matrix
}

func readInputData() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	maybeSize := strings.Split(strings.TrimRight(text, "\n"), " ")
	height, _ := strconv.Atoi(maybeSize[0])
	width, _ := strconv.Atoi(maybeSize[1])
	return width, height
}
