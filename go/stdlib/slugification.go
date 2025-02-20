package main

import (
	"strings"
	"fmt"
	"unicode"
)


func slugify(src string) string {
	var result []rune
	for _, r := range strings.ToLower(src) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			result = append(result, r)
		} else {
			result = append(result, ' ')
		}
	}
	words := strings.Fields(string(result))
	return strings.Join(words, "-")
}
func main() {
    phrase := "Go - Is - Awesome"
    fmt.Println(slugify(phrase))
    // go-is-awesome

    phrase = "Debugging Go code (a status report)"
    fmt.Println(slugify(phrase))
    // tabs-are-all-we-ve-got

	phrase = "Go Talks: \"Cuddle: an App Engine Demo\""
	fmt.Println(slugify(phrase))

	phrase = "Arrays, slices (and strings): The mechanics of 'append'"
	fmt.Println(slugify(phrase))

	phrase = "!Attention, attention!"
	fmt.Println(slugify(phrase))
}