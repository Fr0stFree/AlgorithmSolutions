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
	var builder strings.Builder
	size := len(matrix)

	for rowIdx := 1; rowIdx <= size; rowIdx++ {
		for columnIdx := 1; columnIdx <= size; columnIdx++ {
			var value int
			_, isExist := matrix[rowIdx][columnIdx] 
			if isExist {
				value = 1
			} else {
				value = 0
			}
			builder.WriteString(fmt.Sprintf("%d ", value))
		}
		builder.WriteString("\n")
	}
	fmt.Print(builder.String())
}


func readInputData() map[int]map[int]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	nodesAmount, _ := strconv.Atoi(scanner.Text())

	matrix := make(map[int]map[int]int, nodesAmount)
	for row := range nodesAmount {
		matrix[row+1] = make(map[int]int)
		scanner.Scan()
		edgesAmount, _ := strconv.Atoi(scanner.Text())
		
		for range edgesAmount {
			scanner.Scan()
			toNode, _ := strconv.Atoi(scanner.Text())
			matrix[row+1][toNode] = 1
		}
	}
	return matrix
}
