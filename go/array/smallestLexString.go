// Вам нужно склеить все строчки в одну большую строчку так, чтобы она была лексикографически минимальной.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := readInputData()
	for _, element := range data {
		word := makeSmallestLexString(element)
		fmt.Println(word)
	}
}

func makeSmallestLexString(parts []string) string {
	sort.Slice(parts, func(i, j int) bool {
		return parts[i] + parts[j] < parts[j] + parts[i]
	})
	return strings.Join(parts, "")
}

func readInputData() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	scanner.Scan()
	testAmount, _ := strconv.Atoi(scanner.Text())
	tests := make([][]string, testAmount)
	for testIdx := range testAmount {
		scanner.Scan()
		size, _ := strconv.Atoi(scanner.Text())
		words := make([]string, size)
		for idx := range size {
			scanner.Scan()
			words[idx] = scanner.Text()
		}
		tests[testIdx] = words
	}
	return tests
}
