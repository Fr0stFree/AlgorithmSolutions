package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)


func main() {
	graph := readInputData()
	amount := calcGraphComponents(graph)
	fmt.Println(amount)
}

func calcGraphComponents(graph map[int][]int) int {
	var visitedNodes []int
	var dfs func(node int)

	dfs = func(currentNode int) {
		isVisited := slices.Contains(visitedNodes, currentNode)
		if isVisited {
			return
		}
		visitedNodes = append(visitedNodes, currentNode)
		neighbors := graph[currentNode]
		for _, neighbor := range neighbors {
			dfs(neighbor)
		}
	}
	componentsAmount := 0
	for node := range graph {
		isVisited := slices.Contains(visitedNodes, node)
		if !isVisited {
			componentsAmount++
			dfs(node)
		}
		
	}
	return componentsAmount
}

func readInputData() map[int][]int {
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

	for _ = range edgesAmount {
		scanner.Scan()
		firstNode, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		secondNode, _ := strconv.Atoi(scanner.Text())
		graph[firstNode] = append(graph[firstNode], secondNode)
		graph[secondNode] = append(graph[secondNode], firstNode)

	}

	return graph
}
