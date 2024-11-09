package main

import "fmt"

func main() {
	example := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	modifiers := [][3]int{
		{1, 0, 5},
		{3, 2, 7},
		{5, 3, 4},
	}
	Solve(example, modifiers)
	fmt.Println(example)
}

func Solve(array []int, modifiers [][3]int) {
	//Создадим пустой массив с 0, для хранения разниц
	diffArray := make([]int, len(array))
	// Итерируется по всем вызовам add(x,l,r) и сохраняем модификаторы в diff_array
	for _, modifier := range modifiers {
		x := modifier[0]
		l := modifier[1]
		r := modifier[2]

		diffArray[l] += x
		diffArray[r] -= x
	}

	// Итерируемся по diff_array со вспомогательной переменной adder
	// значение которой складываем со значениями в результирующем массиве
	accumulator := 0
	for index, value := range diffArray {
		accumulator += value
		array[index] += accumulator
	}
}
