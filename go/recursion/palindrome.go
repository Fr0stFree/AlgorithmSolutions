// Создайте рекурсивную функцию isPalindrome, которая проверяет, является ли строка палиндромом.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := readInputStr()
	result := isPalindrome(data)
	if result == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func readInputStr() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text, _ = strings.CutSuffix(text, "\n")
	return text
}

func isPalindrome(word string) bool {
	if len(word) == 0 || len(word) == 1 {
		return true
	}
	if word[0] != word[len(word)-1] {
		return false
	}
	return isPalindrome(word[1 : len(word)-1])
}
