package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func calcDistance(directions []string) int {
	var distance float64
	for _, direction := range directions {
		for _, word := range strings.Fields(direction) {
			if !unicode.IsDigit([]rune(word)[0]) {
				continue
			}
			splitIdx := strings.IndexFunc(word, func(r rune) bool {
				return unicode.IsLetter(r)
			})
			amount, unit := word[:splitIdx], word[splitIdx:]
			parsedAmount, _ := strconv.ParseFloat(amount, 64)

			if unit == "km" {
				parsedAmount *= 1000
			}
			distance += parsedAmount
		}
	}
	return int(distance)
}

func main() {
	directions := []string{
		"straight 1.6km",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	const want = 6000
	got := calcDistance(directions)
	if got != want {
		fmt.Printf("%v: got %v, want %v\n", directions, got, want)
	}
}
