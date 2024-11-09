package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2") == "Fo1r the2 g3ood 4of th5e pe6ople")
}

func Order(sentence string) string {
	splitSentence := strings.Split(sentence, " ")
	result := make([]string, len(splitSentence))

	for _, word := range splitSentence {
		for _, letter := range word {
			if unicode.IsDigit(letter) {
				index, _ := strconv.Atoi(string(letter))
				result[index-1] = word
				break
			}
		}
	}

	return strings.Join(result, " ")
}
