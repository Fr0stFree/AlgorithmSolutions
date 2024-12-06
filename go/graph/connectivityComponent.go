package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)


func main() {
	graph := readInputData()
	components := findGraphComponents(graph)
	printResult(components)
}

func printResult(components [][]int) {
	fmt.Println(len(components))
	sort.Slice(components, func(i, j int) bool {
		return slices.Min(components[i]) < slices.Min(components[j])
	})
	for _, component := range components {
		slices.Sort(component)
		fmt.Printf("%d %s\n", len(component), strings.Trim(fmt.Sprint(component), "[]"))
	}
}

func findGraphComponents(graph map[int][]int) [][]int {
	visitedNodes := make(map[int]bool)
	var dfs func(node int) []int

	dfs = func(currentNode int) []int {
		visitedNodes[currentNode] = true
		neighbors := graph[currentNode]
		localComponents := []int{currentNode}

		for _, neighbor := range neighbors {
			_, isVisited := visitedNodes[neighbor]
			if isVisited {
				continue
			}
			components := dfs(neighbor)
			localComponents = slices.Concat(localComponents, components)
		}
		return localComponents
	}
	components := make([][]int, 0)
	for node := range graph {
		_, isVisited := visitedNodes[node]
		if !isVisited {
			component := dfs(node)
			components = append(components, component)
		}
		
	}
	return components
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