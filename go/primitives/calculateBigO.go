package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var polynomialFactorsPattern = regexp.MustCompile("n(\\^(\\d))?")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	exponent, polynomial := parseExpression(expression)
	result := findBigO(exponent, polynomial)
	if strings.HasPrefix(result, "e") {
		fmt.Println(result)
	} else {
		fmt.Printf("(%s)\n", result)
	}
}

func parseExpression(expression string) (string, string) {
	var exponent, polynomial string
	for _, part := range strings.Split(expression, " * ") {
		if strings.HasPrefix(part, "e") {
			exponent = part
		} else {
			polynomial = part
		}
	}
	return exponent, polynomial
}

func findBigO(exponent, polynomial string) string {
	if exponent != "" {
		return exponent
	}
	matches := polynomialFactorsPattern.FindAllString(polynomial, -1)
	if len(matches) == 0 {
		return "1"
	}
	return matches[0]
}
