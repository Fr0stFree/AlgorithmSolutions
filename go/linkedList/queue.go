// Напишите программу моделирующую работу стека.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Queue[V any] struct {
	elements *list.List
}

func NewQueue[V any]() Queue[V] {
	return Queue[V]{list.New()}
}

func (s *Queue[V]) push(element V) {
	s.elements.PushBack(element)
}
func (s *Queue[V]) pop() V {
	return s.elements.Remove(s.elements.Front()).(V)
}
func (s *Queue[V]) front() V {
	return s.elements.Front().Value.(V)
}
func (s *Queue[V]) size() int {
	return s.elements.Len()
}
func (s *Queue[V]) clear() {
	s.elements.Init()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	queue := NewQueue[int]()
	for {
		scanner.Scan()
		switch scanner.Text() {
		case "push":
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			queue.push(value)
			fmt.Println("ok")
		case "pop":
			if queue.size() == 0 {
				fmt.Println("error")
			} else {
				value := queue.pop()
				fmt.Println(value)
			}
		case "front":
			if queue.size() == 0 {
				fmt.Println("error")
			} else {
				value := queue.front()
				fmt.Println(value)
			}
		case "size":
			value := queue.size()
			fmt.Println(value)
		case "clear":
			queue.clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
