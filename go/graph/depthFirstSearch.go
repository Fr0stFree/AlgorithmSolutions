package main

import "fmt"


func main() {
	graph := makeGraph()
	nodes := traverseDFS(graph, "a")
	fmt.Println(nodes)
}

func traverseDFS(graph map[string][]string, initialNode string) []string {
	var visitedNodes []string
	var dfs func(node string)

	dfs = func(currentNode string) {
		visitedNodes = append(visitedNodes, currentNode)
		neighbors := graph[currentNode]
		for _, neighbor := range neighbors {
			dfs(neighbor)
		}
	}
	dfs(initialNode)
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
