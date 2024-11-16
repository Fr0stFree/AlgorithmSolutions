// Напишите программу, отображающую игровое поле для игры "Сапер".
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	height, width, coords := readInputData()
	field := makeMineField(height, width, coords)
	for _, row := range field {
		fmt.Println(strings.Trim(fmt.Sprint(row), "[]"))
	}

}

func makeMineField(height, width int, coords [][2]int) [][]string {
	field := make([][]string, height)
	for rowIdx := range field {
		field[rowIdx] = make([]string, width)
		for columnIdx := range field[rowIdx] {
			field[rowIdx][columnIdx] = "0"
		}
	}
	for _, coord := range coords {
		field[coord[0]-1][coord[1]-1] = "*"
	}
	for rowIdx, row := range field {
		for columnIdx, value := range row {
			if value != "*" { continue }
			if columnIdx != 0 && field[rowIdx][columnIdx-1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx][columnIdx-1])
				field[rowIdx][columnIdx-1] = fmt.Sprint(value+1)
			}
			if columnIdx != width - 1 && field[rowIdx][columnIdx+1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx][columnIdx+1])
				field[rowIdx][columnIdx+1] = fmt.Sprint(value+1)
			}
			if rowIdx != height - 1 && field[rowIdx+1][columnIdx] != "*" {
				value, _ := strconv.Atoi(field[rowIdx+1][columnIdx])
				field[rowIdx+1][columnIdx] = fmt.Sprint(value+1)
			}
			if rowIdx != 0 && field[rowIdx-1][columnIdx] != "*" {
				value, _ := strconv.Atoi(field[rowIdx-1][columnIdx])
				field[rowIdx-1][columnIdx] = fmt.Sprint(value+1)
			}
			if columnIdx != 0 && rowIdx != 0 && field[rowIdx-1][columnIdx-1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx-1][columnIdx-1])
				field[rowIdx-1][columnIdx-1] = fmt.Sprint(value+1)
			}
			if columnIdx != width - 1 && rowIdx != height - 1 && field[rowIdx+1][columnIdx+1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx+1][columnIdx+1])
				field[rowIdx+1][columnIdx+1] = fmt.Sprint(value+1)
			}
			if columnIdx != width - 1 && rowIdx != 0 && field[rowIdx-1][columnIdx+1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx-1][columnIdx+1])
				field[rowIdx-1][columnIdx+1] = fmt.Sprint(value+1)
			}
			if columnIdx != 0 && rowIdx != height - 1 && field[rowIdx+1][columnIdx-1] != "*" {
				value, _ := strconv.Atoi(field[rowIdx+1][columnIdx-1])
				field[rowIdx+1][columnIdx-1] = fmt.Sprint(value+1)
			}
		}
	}
	return field
}

func readInputData() (int, int, [][2]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	height, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())

	coords := make([][2]int, amount)
	for idx := range amount {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		coords[idx] = [2]int{x, y}
	}

	return height, width, coords
}
