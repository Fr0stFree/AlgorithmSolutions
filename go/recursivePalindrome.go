// Создайте рекурсивную функцию isPalindrome, которая проверяет, является ли строка палиндромом.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := readInput()
	result := isPalindrome(data, 0)
	if result == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func isPalindrome(word string, index int) bool {
	if index >= len(word)/2 {
		return true
	}
	if word[index] != word[len(word)-1-index] {
		return false
	}
	index++
	return isPalindrome(word, index)
}
