package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)


func main() {
	graphs := readInputGraphs()
	for _, graph := range graphs {
		source, destination := graph[0][0], graph[0][1]
		delete(graph, 0)
		distance := calculateDistance(graph, source, destination)
		fmt.Println(distance)
	}

}

func calculateDistance(graph map[int][]int, source, destination int) int {
	visitedNodes := make(map[int]bool)
	queue := list.New()
	queue.PushBack([2]int{source, 0})

	for queue.Len() > 0 {
		element := queue.Back()
		queue.Remove(element)
		value := element.Value.([2]int)
		node, distance := value[0], value[1]
		_, isVisited := visitedNodes[node]
		if isVisited {
			continue
		}
		visitedNodes[node] = true
		
		for _, neighbor := range graph[node] {
			if neighbor == destination {
				return distance + 1
			}
			queue.PushFront([2]int{neighbor, distance+1})
		}
	}
	return -1
}

func readInputGraphs() []map[int][]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	graphsAmount, _ := strconv.Atoi(scanner.Text())
	graphs := make([]map[int][]int, graphsAmount)
	
	for idx := range graphsAmount {
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
		destination, _ := strconv.Atoi(scanner.Text())
		graph[0] = []int{source, destination}
	
		for _ = range edgesAmount {
			scanner.Scan()
			firstNode, _ := strconv.Atoi(scanner.Text())
	
			scanner.Scan()
			secondNode, _ := strconv.Atoi(scanner.Text())
			graph[firstNode] = append(graph[firstNode], secondNode)
		}
		graphs[idx] = graph
	}
	return graphs
}

