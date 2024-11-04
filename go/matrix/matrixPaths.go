// В прямоугольной таблице N×M в начале игрок находится в левой верхней клетке. За один ход ему разрешается перемещаться
// в соседнюю клетку либо вправо, либо вниз (влево и вверх перемещаться запрещено). Посчитайте, сколько есть способов
// у игрока попасть в правую нижнюю клетку.
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
	result := countPaths(width, height)
	fmt.Println(result)

}

func countPaths(width, height int) int {
	matrix := make([][]int, height)
	for idx := 0; idx < height; idx++ {
		matrix[idx] = make([]int, width)
		matrix[idx][0] = 1
	}
	for idx := 0; idx < width; idx++ {
		matrix[0][idx] = 1
	}

	for rowIdx, row := range matrix {
		if rowIdx == 0 {
			continue
		}
		for columnIdx := range row {
			if columnIdx == 0 {
				continue
			}
			amount := matrix[rowIdx-1][columnIdx] + matrix[rowIdx][columnIdx-1]
			matrix[rowIdx][columnIdx] += amount
		}
	}
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

func readInputData() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	maybeSize := strings.Split(strings.TrimRight(text, "\n"), " ")
	width, _ := strconv.Atoi(maybeSize[0])
	height, _ := strconv.Atoi(maybeSize[1])
	return width, height
}
