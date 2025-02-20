package main

import (
	"fmt"
	"regexp"
	"strings"
)

func slugify(src string) string {
	pattern := regexp.MustCompile(`[^a-z0-9-]+`)
	src = strings.ToLower(src)
	src = pattern.ReplaceAllString(src, " ")
	words := strings.Fields(src)
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