package main

import "fmt"

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}

type Map[K comparable, V any] struct {
	storage map[K]V
}

func (m *Map[K, V]) Set(key K, val V) {
	if m.storage == nil {
		m.storage = make(map[K]V)
	}
	m.storage[key] = val
}

func (m *Map[K, V]) Get(key K) V {
	value, _ := m.storage[key]
	return value
}

func (m *Map[K, V]) Keys() []K {
	result := make([]K, len(m.storage))
	index := 0
	for key, _ := range m.storage {
		result[index] = key
		index++
	}
	return result
}

func (m *Map[K, V]) Values() []V {
	result := make([]V, len(m.storage))
	index := 0
	for _, value := range m.storage {
		result[index] = value
		index++
	}
	return result
}
