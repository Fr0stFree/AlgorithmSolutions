package main

import (
	"container/list"
	"fmt"
)


func main() {
	graph := makeGraph()
	nodes := traverseBFS(graph, "a")
	fmt.Println(nodes)
}

func traverseBFS(graph map[string][]string, initialNode string) []string {
	visitedNodes := []string{initialNode}
	queue := list.New()
	queue.PushBack(initialNode)

	for elem := queue.Front(); elem != nil; elem = queue.Front() {
		currentNode := elem.Value.(string)
		queue.Remove(elem)

		neighbors := graph[currentNode]
		for _, neighbor := range neighbors {
			queue.PushBack(neighbor)
			visitedNodes = append(visitedNodes, neighbor)
		}
	}

	return visitedNodes
}

func makeGraph() map[string][]string {
	return map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
}
