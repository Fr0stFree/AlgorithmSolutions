package main

import (
	"fmt"
	"sort"
	"strings"
)

func prettify(m map[string]int) string {
	keys := make([]string, len(m))
	idx := 0
	for key, _ := range m {
		keys[idx] = key
		idx++
	}
	sort.Strings(keys)

	var offset = "    "
	if len(keys) <= 1 {
		offset = " "
	}
	var builder strings.Builder
	builder.WriteRune('{')
	if len(keys) > 1 {
		builder.WriteRune('\n')
	}
	for _, key := range keys {
		value := m[key]
		builder.WriteString(fmt.Sprintf("%s%s: %d", offset, key, value))
		if len(keys) > 1 {
			builder.WriteString(",\n")
		} else {
			builder.WriteString(offset)
		}
	}
	builder.WriteRune('}')
	
	result := builder.String()
	return result
}


func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	prettify(m)

	fmt.Println()
	m = map[string]int{"one": 1}
	prettify(m)

	fmt.Println()
	m = map[string]int{}
	prettify(m)
}
