package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	graphs := readInputData()
	for _, graph := range graphs {
		if isTree(graph) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func isTree(graph map[int][]int) bool {
	visitedNodes := make(map[int]bool)
	previousNodes := make(map[int]bool)
	var isGraphContainCycle func (int) bool

	isGraphContainCycle = func(currentNode int) bool {
		if visitedNodes[currentNode] {
			return true
		}
		visitedNodes[currentNode] = true
		previousNodes[currentNode] = true
		for _, node := range graph[currentNode] {
			if !previousNodes[node] && isGraphContainCycle(node) {
				return true
			}
		}
		previousNodes[currentNode] = false
		return false
	}
	
	hasCycle := isGraphContainCycle(1)
	hasOnlyOneComponent := len(visitedNodes) == len(graph)
	return !hasCycle && hasOnlyOneComponent
}

func readInputData() []map[int][]int {
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
	
		for _ = range edgesAmount {
			scanner.Scan()
			firstNode, _ := strconv.Atoi(scanner.Text())
	
			scanner.Scan()
			secondNode, _ := strconv.Atoi(scanner.Text())
			graph[firstNode] = append(graph[firstNode], secondNode)
			graph[secondNode] = append(graph[secondNode], firstNode)
	
		}
		graphs[idx] = graph
	}
	return graphs
}

