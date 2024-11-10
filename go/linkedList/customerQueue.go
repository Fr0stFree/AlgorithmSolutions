// Написать программу, которая моделирует работу в банке. Клиенты банка приходят и становятся в очередь.
// Каждому клиенту присваивается уникальный идентификатор ID. Когда клиент обслуживается, он уходит из очереди.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	queue := list.New()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	operationsAmount, _ := strconv.Atoi(scanner.Text())
	for _ = range operationsAmount {
		scanner.Scan()
		switch scanner.Text() {
		case "ADD":
			scanner.Scan()
			customerId, _ := strconv.Atoi(scanner.Text())
			queue.PushFront(customerId)
		case "NEXT":
			next := queue.Back()
			if next == nil {
				fmt.Println("Queue is empty")
			} else {
				queue.Remove(next)
			}
		case "COUNT":
			fmt.Println(queue.Len())
		}
	}
}
