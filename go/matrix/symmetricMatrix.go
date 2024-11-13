// Проверьте, является ли двумерный массив симметричным относительно главной диагонали.
// Главная диагональ — та, которая идёт из левого верхнего угла двумерного массива в правый нижний.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInputData()
	result := isMatrixSymmetric(data)
	if result {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func isMatrixSymmetric(matrix [][]int) bool {
	for rowIdx, row := range matrix {
		for columnIdx := range row {
			if matrix[rowIdx][columnIdx] != matrix[columnIdx][rowIdx] {
				return false
			}
		}
	}
	return true
}

func readInputData() [][]int {
	reader := bufio.NewReader(os.Stdin)
	maybeSize, _ := reader.ReadString('\n')
	maybeSize = strings.TrimRight(maybeSize, "\n")
	size, _ := strconv.Atoi(maybeSize)

	result := make([][]int, size)
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		text, _ := reader.ReadString('\n')
		maybeNumbers := strings.Split(strings.TrimRight(text, "\n"), " ")
		row := make([]int, len(maybeNumbers))
		for index, elem := range maybeNumbers {
			row[index], _ = strconv.Atoi(elem)
		}
		result[rowIdx] = row
	}

	return result
}
