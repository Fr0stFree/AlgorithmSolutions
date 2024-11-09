// Найдите срединный элемент в односвязном списке. Если список содержит четное количество элементов,
// верните второй элемент из двух средних.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"slices"
)

func main() {
	sequence := readInputData()
	isCorrect := isSequenceCorrect(sequence)
	if isCorrect {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func isSequenceCorrect(sequence string) bool {
	stack := list.New()
	openBrackets := []string{"{", "[", "("}
	closeBrackets := []string{"}", "]", ")"}
	for _, bracket := range sequence {
		if slices.Contains(openBrackets, string(bracket)) {
			stack.PushBack(string(bracket))
		} else {
			if stack.Back() == nil {
				return false
			}
			prev := stack.Remove(stack.Back()).(string)
			prevIdx := slices.Index(openBrackets, prev)
			if prevIdx == -1 || slices.Index(closeBrackets, string(bracket)) != prevIdx {
				return false
			}
		}
	}
	return stack.Len() == 0
}

func readInputData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}
