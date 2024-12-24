package main

import (
	"sort"
	"strings"
)

// JoinWords объединяет слова из двух строк, приводит к нижнему регистру,
// сортирует, удаляет дубликаты и возвращает получившиеся слова.
func JoinWords(first, second string) []string {
	words1 := split(first)
	words2 := split(second)
	words := join(words1, words2)
	words = lower(words)
	words = sorted(words)
	words = uniq(words)
	return words
}

// split разбивает строку на слова.
func split(str string) []string {
	return strings.Fields(str)
}

// join объединяет два списка слов.
func join(words1, words2 []string) []string {
	words := make([]string, len(words1)+len(words2))
	idx := 0
	for _, word := range words1 {
		words[idx] = word
		idx++
	}
	for _, word := range words2 {
		words[idx] = word
		idx++
	}
	return words
}

// lower приводит слова к нижнему регистру.
func lower(words []string) []string {
	for idx, word := range words {
		words[idx] = strings.ToLower(word)
	}
	return words
}

// sorted сортирует слова.
func sorted(words []string) []string {
	sort.Strings(words)
	return words
}

// uniq возвращает уникальные слова.
// Исходный список слов должен быть отсортирован.
func uniq(words []string) []string {
	uniq := []string{}
	current := ""
	for _, word := range words {
		if word == current {
			continue
		}
		if current != "" {
			uniq = append(uniq, current)
		}
		current = word
	}
	if current != "" {
		uniq = append(uniq, current)
	}
	return uniq
}