package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var builder strings.Builder

	for scanner.Scan() {
		word := []rune(scanner.Text())
		builder.WriteRune(unicode.ToUpper(word[0]))
		for _, r := range word[1:len(word)] {
			builder.WriteRune(unicode.ToLower(r))
		}
		builder.WriteRune(' ')
	}
	err := scanner.Err(); if err != nil {
		panic(err)
	}

	fmt.Println(builder.String())
}