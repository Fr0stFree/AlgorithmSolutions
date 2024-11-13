// Дано число n - размер квадратной матрицы. Создайте двумерный массив и заполните его нулями, единицами и двойками по
// правилу: единицы располагаются на побочной диагонали, нули над единицами, а двойки под. Выведите полученный двумерный массив на экран.
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
	result := fillMatrix(data)
	for row := range result {
		fmt.Println(strings.Trim(fmt.Sprint(result[row]), "[]"))
	}

}

func fillMatrix(size int) [][]int {
	result := make([][]int, size)

	for rowIdx := range size {
		result[rowIdx] = make([]int, size)
		for columnIdx := range size {
			var value int
			if rowIdx+columnIdx > size-1 {
				value = 2
			} else if rowIdx+columnIdx < size-1 {
				value = 0
			} else {
				value = 1
			}
			result[rowIdx][columnIdx] = value
		}
	}

	return result
}

func readInputData() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	amount, _ := strconv.Atoi(text)
	return amount
}
