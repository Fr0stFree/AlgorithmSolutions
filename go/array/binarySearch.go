// У вас есть отсортированный по названию список книг в библиотеке и вы ищете конкретную книгу.
// Ваша задача - написать программу, которая с помощью бинарного поиска определит, есть ли книга в списке.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	books, lookingBook := readInputData()
	isExists := searchBook(books, lookingBook)
	if isExists {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func searchBook(books []string, lookingBook string) bool {
	leftIdx, rightIdx := 0, len(books) - 1
	for leftIdx <= rightIdx {
		middleIdx := (rightIdx+leftIdx)/2
		middleBook := books[middleIdx]
		switch strings.Compare(lookingBook, middleBook) {
		case 0:
			return true
		case 1:
			leftIdx = middleIdx + 1
		case -1:
			rightIdx = middleIdx - 1
		}
	}
	return false
}

func readInputData() ([]string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	books := make([]string, size)
	for index := range size {
		scanner.Scan()
		books[index] = scanner.Text()
	}
	scanner.Scan()
	lookingBook := scanner.Text()
	return books, lookingBook
}
