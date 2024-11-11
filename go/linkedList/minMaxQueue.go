package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
)

type MaxStack struct {
	stack    *list.List
	maxStack *list.List
}

func (s *MaxStack) Push(value int) {
	s.stack.PushBack(value)
	maxElement := s.maxStack.Back()
	if maxElement == nil || s.maxStack.Back().Value.(int) <= value {
		s.maxStack.PushBack(value)
	}
}

func (s *MaxStack) Pop() int {
	value := s.stack.Remove(s.stack.Back()).(int)
	maxElement := s.maxStack.Back()
	if maxElement != nil && maxElement.Value.(int) == value {
		s.maxStack.Remove(maxElement)
	}
	return value
}

func (s *MaxStack) Size() int {
	return s.stack.Len()
}

func (s *MaxStack) Max() int {
	maxElement := s.maxStack.Back()
	if maxElement == nil {
		return math.MinInt
	}
	return maxElement.Value.(int)
}

func NewMaxStack() *MaxStack {
	return &MaxStack{stack: list.New(), maxStack: list.New()}
}

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

func (s *MinStack) Push(value int) {
	s.stack.PushBack(value)
	minElement := s.minStack.Back()
	if minElement == nil || s.minStack.Back().Value.(int) >= value {
		s.minStack.PushBack(value)
	}
}

func (s *MinStack) Pop() int {
	value := s.stack.Remove(s.stack.Back()).(int)
	minElem := s.minStack.Back()
	if minElem != nil && minElem.Value.(int) == value {
		s.minStack.Remove(minElem)
	}
	return value
}

func (s *MinStack) Size() int {
	return s.stack.Len()
}

func (s *MinStack) Min() int {
	minElem := s.minStack.Back()
	if minElem == nil {
		return math.MaxInt
	}
	return minElem.Value.(int)
}

func NewMinStack() *MinStack {
	return &MinStack{stack: list.New(), minStack: list.New()}
}

type MaxQueue struct {
	backStack  *MaxStack
	frontStack *MaxStack
}

func (s *MaxQueue) Enqueue(value int) {
	s.backStack.Push(value)
}

func (s *MaxQueue) Dequeue() int {
	if s.frontStack.Size() == 0 {
		for s.backStack.Size() != 0 {
			s.frontStack.Push(s.backStack.Pop())
		}
	}
	return s.frontStack.Pop()
}

func (s *MaxQueue) Max() int {
	return max(s.backStack.Max(), s.frontStack.Max())
}

func (s *MaxQueue) Size() int {
	return s.backStack.Size() + s.frontStack.Size()
}

func NewMaxQueue() *MaxQueue {
	return &MaxQueue{NewMaxStack(), NewMaxStack()}
}

type MinQueue struct {
	backStack  *MinStack
	frontStack *MinStack
}

func (s *MinQueue) Enqueue(value int) {
	s.backStack.Push(value)
}

func (s *MinQueue) Dequeue() int {
	if s.frontStack.Size() == 0 {
		for s.backStack.Size() != 0 {
			s.frontStack.Push(s.backStack.Pop())
		}
	}
	return s.frontStack.Pop()
}

func (s *MinQueue) Min() int {
	return min(s.backStack.Min(), s.frontStack.Min())
}

func (s *MinQueue) Size() int {
	return s.backStack.Size() + s.frontStack.Size()
}

func NewMinQueue() *MinQueue {
	return &MinQueue{NewMinStack(), NewMinStack()}
}

type MinMaxQueue struct {
	minQueue *MinQueue
	maxQueue *MaxQueue
}

func (q *MinMaxQueue) Enqueue(value int) {
	q.maxQueue.Enqueue(value)
	q.minQueue.Enqueue(value)
}

func (q *MinMaxQueue) Dequeue() int {
	q.maxQueue.Dequeue()
	return q.minQueue.Dequeue()
}

func (q *MinMaxQueue) Size() int {
	return q.minQueue.Size()
}

func (q *MinMaxQueue) Min() int {
	return q.minQueue.Min()
}

func (q *MinMaxQueue) Max() int {
	return q.maxQueue.Max()
}

func NewMinMaxQueue() *MinMaxQueue {
	return &MinMaxQueue{NewMinQueue(), NewMaxQueue()}
}

func main() {
	potential, values := readInputData()
	result := calcBestPotential(potential, values)
	fmt.Println(result)
}

func calcBestPotential(potentialThreshold int, values []int) int {
	queue := NewMinMaxQueue()
	bestLength := len(values) + 1
	leftIdx := 0

	for rightIdx, value := range values {
		queue.Enqueue(value)

		for queue.Max()-queue.Min() > potentialThreshold {
			bestLength = min(bestLength, rightIdx-leftIdx+1)
			queue.Dequeue()
			leftIdx++
		}
	}
	if bestLength == len(values)+1 {
		return 0
	}
	return bestLength
}

func readInputData() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	potential, _ := strconv.Atoi(scanner.Text())
	values := make([]int, amount)
	for idx := range amount {
		scanner.Scan()
		values[idx], _ = strconv.Atoi(scanner.Text())
	}
	return potential, values
}
