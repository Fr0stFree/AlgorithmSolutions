package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	expected := "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S"
	result := Vaporcode("Lets go to the movies")
	if result != expected {
		fmt.Printf("Expected: '%s'\nGot: '%s'", expected, result)
	} else {
		fmt.Println("PASSED")
	}
}

func Vaporcode(s string) string {
	letters := make([]string, len([]rune(s)))

	for index, letter := range s {
		if unicode.IsSpace(letter) {
			continue
		}
		letters[index] = fmt.Sprintf("%c  ", unicode.ToUpper(letter))
	}
	sentence := strings.Trim(strings.Join(letters, ""), " ")
	return sentence
}
