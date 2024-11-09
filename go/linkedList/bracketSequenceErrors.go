// Дана строка, составленная из круглых скобок. Определите, какое наименьшее количество символов необходимо удалить
// из этой строки, чтобы оставшиеся символы образовывали правильную скобочную последовательность.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	sequence := readInputData()
	errorAmount := calcErrorAmount(sequence)
	fmt.Println(errorAmount)
}

func calcErrorAmount(sequence string) int {
	var errorCounter int
	stack := list.New()
	for _, char := range sequence {
		letter := string(char)
		if letter == "(" {
			stack.PushBack(letter)
		} else {
			if stack.Back() == nil {
				errorCounter++
				continue
			}
			stack.Remove(stack.Back())
		}
	}
	return errorCounter + stack.Len()
}

func readInputData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}
