package main

import (
	"fmt"
	"os"
	"bufio"
)

func readLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sentences := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return sentences, nil
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
