// n детей собрались на большой поляне, встали в круг, и решили познакомиться.
// У каждого ребёнка есть имя. В определенном порядке каждый ребёнок кричит меня зовут X, слева от меня стоит L,
// справа от меня стоит R и выходит из круга. Когда в круге остаётся 3 человека, они перестают кричать и спокойно расходятся.
// Известны имена детей и порядок, в котором они выходили из круга. Восстановите, что они кричали.
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Child struct {
	index int
	name  string
}

func main() {
	names, numbers := readInputData()
	children := initList(names)
	mutateList(children, numbers)
}

func mutateList(children *list.List, numbers []int) {
	elem := children.Front()

	for _, number := range numbers {
		for {
			child := elem.Value.(Child)
			if child.index == number {
				prev := elem.Prev()
				if prev == nil {
					prev = children.Back()
				}
				next := elem.Next()
				if next == nil {
					next = children.Front()
				}
				fmt.Printf("%s %s\n", prev.Value.(Child).name, next.Value.(Child).name)
				toRemove := elem
				elem = elem.Prev()
				children.Remove(toRemove)
				break
			} else if child.index < number {
				elem = elem.Next()
			} else {
				elem = elem.Prev()
			}
		}
	}
}

func initList(values []string) *list.List {
	elements := list.New()
	for index, value := range values {
		elements.PushBack(Child{index + 1, value})
	}
	return elements
}

func readInputData() ([]string, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())

	names := make([]string, amount)
	for idx := range amount {
		scanner.Scan()
		names[idx] = scanner.Text()
	}

	numbers := make([]int, amount-3)
	for idx := range amount - 3 {
		scanner.Scan()
		numbers[idx], _ = strconv.Atoi(scanner.Text())
	}
	return names, numbers
}
