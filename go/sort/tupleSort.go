package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tuple struct {
	element [2]int
	index int
}
func main() {
	data := readInputData()
	result := sortTuples(data)

	var builder strings.Builder

	for idx, value := range result {
		if idx != len(result) - 1 {
			builder.WriteString(fmt.Sprintf("%d ", value.index))
		} else {
			builder.WriteString(fmt.Sprintf("%d", value.index))
		}
		
	}
	fmt.Println(builder.String())
}

func sortTuples(elements [][2]int) []Tuple {
	result := make([]Tuple, len(elements))
	
	for idx, element := range elements {
		result[idx] = Tuple{element, idx+1}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].element[0] != result[j].element[0] {
			return result[i].element[0] < result[j].element[0]
		}
		if result[i].element[1] != result[j].element[1] {
			return result[i].element[1] < result[j].element[1]
		}
		return result[i].index < result[j].index
	})
	return result
}

func readInputData() [][2]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	result := make([][2]int, size)
	for idx := range size {
		scanner.Scan()
		result[idx][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		result[idx][1], _ = strconv.Atoi(scanner.Text())
	}
	return result
}
