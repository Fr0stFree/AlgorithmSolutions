package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
)


func main() {
	graph, source, maxDistance := readInputGraph()
	amountOfAdjacentNodes := calcAdjacentNodes(graph, source, maxDistance)
	fmt.Println(amountOfAdjacentNodes)
}

func calcAdjacentNodes(graph map[int][]int, source, maxDistance int) int {
	destinations := make([]int, len(graph))
	for idx := range destinations {
		destinations[idx] = math.MaxInt
	}

	visitedNodes := make(map[int]bool)
	queue := list.New()
	queue.PushBack([2]int{source, 0})

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		value := element.Value.([2]int)
		node, destination := value[0], value[1]
		
		if _, ok := visitedNodes[node]; ok {
			continue
		}
		visitedNodes[node] = true
		destinations[node-1] = destination

		for _, neighbor := range graph[node] {
			if _, ok := visitedNodes[neighbor]; !ok {
				queue.PushBack([2]int{neighbor, destination+1})
			}
		}

	}
	amountOfAdjacentNodes := 0
	for _, destination := range destinations {
		if destination <= maxDistance && destination != 0 {
			amountOfAdjacentNodes++
		}
	}
	return amountOfAdjacentNodes
}

func readInputGraph() (map[int][]int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	nodesAmount, _ := strconv.Atoi(scanner.Text())
	graph := make(map[int][]int, nodesAmount)
	for idx := range nodesAmount {
		graph[idx+1] = []int{}
	}
	scanner.Scan()
	edgesAmount, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	source, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	maxDistance, _ := strconv.Atoi(scanner.Text())

	for _ = range edgesAmount {
		scanner.Scan()
		firstNode, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		secondNode, _ := strconv.Atoi(scanner.Text())
		graph[firstNode] = append(graph[firstNode], secondNode)	
	}

	return graph, source, maxDistance
}

