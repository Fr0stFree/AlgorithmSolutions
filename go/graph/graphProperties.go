package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	graph := readInputData()
	hinges, isolations, valences := calcGraphProperties(graph)
	fmt.Println(hinges)
	fmt.Println(isolations)
	fmt.Println(strings.Trim(fmt.Sprint(valences), "[]"))
}

func calcGraphProperties(graph map[int]map[int]int) (int, int, []int) {
	var hingesAmount, isolatedAmount int
	var maxValences []int
	valences := make(map[int]int)
	for node, connections := range graph {
		if connections[node] == 1 {
			hingesAmount++
		}

		isIsolated := true
		for neighbor, isConnected := range connections {
			if isConnected == 1 && neighbor != node {
				isIsolated = false
				break
			}
		}
		if isIsolated {
			isolatedAmount++
		}
		
		valence := 0
		for _, isConnected := range connections {
			valence += isConnected
		}
		valences[node] = valence
	}
	maxValence := 0
	for _, valence := range valences {
		if valence > maxValence {
			maxValence = valence
		}
	}
	for node, valence := range valences {
		if valence == maxValence {
			maxValences = append(maxValences, node)
		}
	}
	return hingesAmount, isolatedAmount, maxValences
}

func readInputData() map[int]map[int]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	nodesAmount, _ := strconv.Atoi(scanner.Text())

	graph := make(map[int]map[int]int, nodesAmount)
	for idx := range nodesAmount {
		graph[idx+1] = make(map[int]int)
		for innerIdx := range nodesAmount {
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			graph[idx+1][innerIdx+1] = value
		}
	}

	return graph
}
