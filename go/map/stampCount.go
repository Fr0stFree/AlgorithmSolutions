package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const totalAmount = 15_000
	stamps := readInput()
	amount := countStamps(stamps)
	fmt.Println(totalAmount-amount)
}


func countStamps(stamps []int) int {
	unique := make(map[int]bool)
	for _, stamp := range stamps {
		unique[stamp] = true
	}
	return len(unique)
}


func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	values := make([]int, size)
	for idx := range size {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return values
}
