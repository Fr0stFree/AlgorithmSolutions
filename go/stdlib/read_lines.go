package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	sentences := strings.Split(string(data), "\n")
	return sentences[:len(sentences)-1], nil
}

func main() {
	lines, err := readLines("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
