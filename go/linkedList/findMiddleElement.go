// Найдите срединный элемент в односвязном списке. Если список содержит четное количество элементов,
// верните второй элемент из двух средних.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := readInputData()
	elements := initList(&numbers)
	middle := findMiddle(elements)
	fmt.Println(middle.Value)
}

func findMiddle(elements *list.List) *list.Element {
	for elem, idx := elements.Front(), 0; elem != nil; elem, idx = elem.Next(), idx+1 {
		if idx == elements.Len()/2 {
			return elem
		}
	}
	return nil
}

func initList(numbers *[]int) *list.List {
	elements := list.New()
	for _, element := range *numbers {
		elements.PushBack(element)
	}
	return elements
}

func readInputData() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())

	numbers := make([]int, amount)
	for idx := range amount {
		scanner.Scan()
		number, _ := strconv.Atoi(scanner.Text())
		numbers[idx] = number
	}
	return numbers
}
