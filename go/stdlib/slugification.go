package main

import (
	"strings"
	"fmt"
)

const ALLOWED_LETTERS = "-1234567890abcdefghijklmnopqrstuvwxyz"
const SPECIAL_CHARACTERS = "!#$%^&*+./_ "

func slugify(src string) string {
	src = strings.ToLower(src)
	var builder strings.Builder 

	for _, letter := range src {
		if strings.Contains(SPECIAL_CHARACTERS, string(letter)) {
			builder.WriteRune('-')
			continue
		}
		if strings.Contains(ALLOWED_LETTERS, string(letter)) {
			builder.WriteRune(letter)
			continue
		}
	}
    return strings.Trim(builder.String(), " -")
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