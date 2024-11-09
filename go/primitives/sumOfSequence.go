package main

func main() {
	println(SequenceSum(1, 5, 1) == 15)
}

func SequenceSum(start, end, step int) int {
	var sum int
	for index := start; index <= end; index += step {
		sum += index
	}
	return sum
}
