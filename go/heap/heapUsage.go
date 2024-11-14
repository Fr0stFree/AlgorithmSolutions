// Изначально у вас есть пустое мультимножество.
// Нужно обработать несколько запросов двух видов:
// 1. «добавить целое число в мультимножество» и
// 2. «извлечь из мультимножества минимальное по значению число».
// Найдите срединный элемент в односвязном списке. Если список содержит четное количество элементов,
// верните второй элемент из двух средних.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"container/heap"
)

type IntHeap []int
    
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) 		 { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	commands := newCommandReader()
	h := &IntHeap{}
	heap.Init(h)

	for command := range commands {
		switch command.(type) {
		case PopCommand:
			value := heap.Pop(h)
			fmt.Println(value)
		case AddCommand:
			adder := command.(AddCommand).adder
			heap.Push(h, adder)
		}
	}
}

type Command interface {}

type AddCommand struct {
	adder int
}

type PopCommand struct {}

func newCommandReader() chan Command {
	commands := make(chan Command)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		scanner.Scan()
		size, _ := strconv.Atoi(scanner.Text())
		for range size {
			scanner.Scan()
			switch scanner.Text() {
			case "+":
				scanner.Scan()
				adder, _ := strconv.Atoi(scanner.Text())
				commands <- AddCommand{adder}
			case "-":
				commands <- PopCommand{}
			default:
				panic("unknown command")
			}
		}
		close(commands)
	}()

	return commands
}
