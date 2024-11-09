package main

import "fmt"

func main() {
	fmt.Println(QueueTime([]int{5, 3, 4}, 1) == 12)
	fmt.Println(QueueTime([]int{10, 2, 3, 3}, 2) == 10)
	fmt.Println(QueueTime([]int{1, 2, 3, 4, 5}, 100) == 5)
}

func QueueTime(customers []int, n int) int {
	var timePassed int
	servedCustomers := make([]int, n)

	for {
		if isFinished(append(servedCustomers, customers...)) {
			break
		}
		for index := range servedCustomers {
			customer := servedCustomers[index]
			if customer <= 0 {
				if len(customers) == 0 {
					continue
				}
				customer, customers = customers[0], customers[1:]
			}
			servedCustomers[index] = customer - 1
		}
		timePassed++
	}
	return timePassed
}

func isFinished(queue []int) bool {
	for _, customer := range queue {
		if customer != 0 {
			return false
		}
	}
	return true
}
