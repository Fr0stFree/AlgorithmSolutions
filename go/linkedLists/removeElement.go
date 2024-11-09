// Удалите из односвязного списка все узлы с заданным значением.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, lookingElement := readInputData()
	elements := initList(&numbers)
	dropElement(elements, lookingElement)
	printElements(elements)
}

func printElements(elements *list.List) {
	if elements.Len() == 0 {
		fmt.Println("None")
	}
	for elem := elements.Front(); elem != nil; elem = elem.Next() {
		fmt.Printf("%-2d", elem.Value)
	}
	fmt.Printf("\n")
}

func dropElement(elements *list.List, lookingElem int) {
	var toDelete []*list.Element
	for elem := elements.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == lookingElem {
			toDelete = append(toDelete, elem)
		}
	}
	for _, elem := range toDelete {
		elements.Remove(elem)
	}
}

func initList(numbers *[]int) *list.List {
	elements := list.New()
	for _, element := range *numbers {
		elements.PushBack(element)
	}
	return elements
}

func readInputData() ([]int, int) {
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

	scanner.Scan()
	lookingElem, _ := strconv.Atoi(scanner.Text())
	return numbers, lookingElem
}
