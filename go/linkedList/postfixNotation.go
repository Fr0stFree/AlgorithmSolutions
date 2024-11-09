// Вычислите выражение в постфиксной записи
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data := readInputData()
	result := calculatePostfixNotation(data)
	fmt.Println(result)
}

func calculatePostfixNotation(elements []string) string {
	stack := list.New()

	for _, elem := range elements {
		if isStrInt(elem) {
			stack.PushBack(elem)
		} else {
			second, _ := strconv.Atoi(stack.Remove(stack.Back()).(string))
			first, _ := strconv.Atoi(stack.Remove(stack.Back()).(string))
			result := strconv.Itoa(calculateExpression(first, second, elem))
			stack.PushBack(result)
		}
	}
	return stack.Front().Value.(string)
}

func calculateExpression(first, second int, operand string) int {
	switch operand {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		panic("unknown operator")
	}

}

func isStrInt(s string) bool {
	for _, value := range s {
		if !unicode.IsDigit(value) {
			return false
		}
	}
	return true
}

func readInputData() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	var elements []string
	for _, elem := range strings.Split(scanner.Text(), " ") {
		elements = append(elements, elem)
	}
	return elements
}
