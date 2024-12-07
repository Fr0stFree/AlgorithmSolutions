package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var builder strings.Builder
	matrix, questions := readInputData()
	for _, question := range questions {
		source, destination := question[0], question[1]
		_, isExist := matrix[source][destination]
		if isExist {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}
	fmt.Println(builder.String())
}


func readInputData() (map[int]map[int]bool, [][2]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	nodesAmount, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	edgesAmount, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	questionsAmount, _ := strconv.Atoi(scanner.Text())

	matrix := make(map[int]map[int]bool, nodesAmount)
	for idx := range nodesAmount {
		matrix[idx+1] = make(map[int]bool)
	}
	for range edgesAmount {
		scanner.Scan()
		firstNode, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		secondNode, _ := strconv.Atoi(scanner.Text())
		
		matrix[firstNode][secondNode] = true
		matrix[secondNode][firstNode] = true
	}
	questions := make([][2]int, questionsAmount)
	for idx := range questionsAmount {
		scanner.Scan()

		scanner.Scan()
		source, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		destination, _ := strconv.Atoi(scanner.Text())
		questions[idx] = [2]int{source, destination}

	}
	return matrix, questions
}
