package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	first, second := readInputData()
	winner := calcWinner(first, second)
	fmt.Println(winner)
}

func calcWinner(first, second *list.List) string {
	for limiter := 1; limiter < 1_000_000; limiter++ {
		firstCard := first.Remove(first.Front()).(int)
		secondCard := second.Remove(second.Front()).(int)

		if secondCard > firstCard || (firstCard == 9 && secondCard == 0) {
			second.PushBack(firstCard)
			second.PushBack(secondCard)
		} else {
			first.PushBack(firstCard)
			first.PushBack(secondCard)
		}
		if first.Len() == 0 {
			return fmt.Sprintf("second %d", limiter)
		}
		if second.Len() == 0 {
			return fmt.Sprintf("first %d", limiter)
		}

	}
	return "botva"
}

func readInputData() (*list.List, *list.List) {
	first, second := list.New(), list.New()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	for _, char := range strings.Split(scanner.Text(), " ") {
		value, _ := strconv.Atoi(char)
		first.PushBack(value)
	}

	scanner.Scan()
	for _, char := range strings.Split(scanner.Text(), " ") {
		value, _ := strconv.Atoi(char)
		second.PushBack(value)
	}
	return first, second
}
