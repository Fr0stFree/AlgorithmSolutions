// Дана строка, содержащая только символы '('. Определите максимальную глубину вложенности скобок в данном выражении.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

func main() {
	sequence := readInputData()
	depth := calcDepth(sequence)
	fmt.Println(depth)
}

func calcDepth(sequence string) int {
	deque := list.New()
	maxDepth := math.MinInt
	var depth int

	for _, bracket := range sequence {
		if string(bracket) == "(" {
			deque.PushBack(string(bracket))
			depth++
		} else {
			deque.Remove(deque.Back())
			maxDepth = max(depth, maxDepth)
			depth--
		}
	}
	return maxDepth
}

func readInputData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}
