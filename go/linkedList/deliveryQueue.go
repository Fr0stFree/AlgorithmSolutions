// Допустим, вы управляете центром обработки заявок. Заявки приходят через некоторые промежутки времени. Каждая заявка
// требует определенное время для ее обработки. Ваша задача - написать программу, которая определит время начала
// обработки каждой заявки, учитывая, что центр может обрабатывать только одну заявку за раз.
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
	data := readInputData()
	result := calcDeliveryTime(data)
	for _, value := range result {
		fmt.Println(value)
	}
}

func calcDeliveryTime(deliveries [][2]int) []int {
	result := make([]int, len(deliveries))
	queue := list.New()
	currentTime := deliveries[0][0]

	for idx, delivery := range deliveries {
		queue.PushBack(delivery)
		for queue.Len() > 0 {
			first := queue.Front().Value.([2]int)
			if first[0] <= currentTime {
				result[idx] = currentTime
				currentTime += first[1]
				queue.Remove(queue.Front())
			} else {
				break
			}
		}
	}

	return result
}

func readInputData() [][2]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())

	result := make([][2]int, amount)
	for idx := range amount {
		scanner.Scan()
		numbers := strings.Split(scanner.Text(), " ")
		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])
		result[idx] = [2]int{first, second}
	}
	return result
}
