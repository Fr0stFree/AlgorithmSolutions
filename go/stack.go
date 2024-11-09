// Напишите программу моделирующую работу стека.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Stack[V comparable] struct {
	elements *list.List
}

func NewStack[V comparable]() Stack[V] {
	return Stack[V]{list.New()}
}

func (s *Stack[V]) push(element V) {
	s.elements.PushBack(element)
}
func (s *Stack[V]) pop() V {
	return s.elements.Remove(s.elements.Back()).(V)
}
func (s *Stack[V]) back() V {
	return s.elements.Back().Value.(V)
}
func (s *Stack[V]) size() int {
	return s.elements.Len()
}
func (s *Stack[V]) clear() {
	s.elements.Init()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	stack := NewStack[int]()
	for {
		scanner.Scan()
		switch scanner.Text() {
		case "push":
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			stack.push(value)
			fmt.Println("ok")
		case "pop":
			value := stack.pop()
			fmt.Println(value)
		case "back":
			value := stack.back()
			fmt.Println(value)
		case "size":
			value := stack.size()
			fmt.Println(value)
		case "clear":
			stack.clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
