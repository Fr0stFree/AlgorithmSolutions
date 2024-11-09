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
	var builder strings.Builder

	for _, letter := range s {
		if unicode.IsSpace(letter) {
			continue
		}
		builder.WriteRune(unicode.ToUpper(letter))
		builder.WriteString("  ")
	}

	return strings.Trim(builder.String(), " ")
}
