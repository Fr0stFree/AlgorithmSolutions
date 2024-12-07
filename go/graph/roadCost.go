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
	graphs := readInputGraphs()
	for _, graph := range graphs {
		costs := calcCostsToCapital(graph)
		fmt.Println(strings.Trim(fmt.Sprint(costs), "[]"))
	}

}

func calcCostsToCapital(graph map[int][]int) []int {
	destinations := make([]int, len(graph))
	for idx := range destinations {
		destinations[idx] = -1
	}
	reversedGraph := make(map[int][]int)
	for from, neighbors := range graph {
		for _, to := range neighbors {
			reversedGraph[to] = append(reversedGraph[to], from)
		}
	}

	visitedNodes := make(map[int]bool)
	queue := list.New()
	queue.PushBack([2]int{1, 0})

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

		for _, neighbor := range reversedGraph[node] {
			if _, ok := visitedNodes[neighbor]; !ok {
				queue.PushBack([2]int{neighbor, destination+1})
			}
		}

	}

	return destinations
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

