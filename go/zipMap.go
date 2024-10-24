package main

import "fmt"

func main() {
	keys := []string{"one", "two", "thr"}
	vals := []int{11, 22, 33}

	m := ZipMap(keys, vals)
	fmt.Println(m)
}

func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	mapSize := min(len(keys), len(vals))
	result := make(map[K]V, mapSize)
	for index := 0; index < mapSize; index++ {
		result[keys[index]] = vals[index]
	}
	return result
}
