package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestJoinWords(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   []string
	}{
		{"hello", "world", []string{"hello", "world"}},
		{"world", "hello", []string{"hello", "world"}},
		{"", "world", []string{"world"}},
		{"hello", "", []string{"hello"}},
		{"", "", []string{}},
		{"Python is", "awesome", []string{"awesome", "is", "python"}},
		{"Python is awesome", "PHP is not awesome",
			[]string{"awesome", "is", "not", "php", "python"}},
	}
	for _, test := range tests {
		got := JoinWords(test.s1, test.s2)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("'%s' + '%s': got %v, want %v", test.s1, test.s2, got, test.want)
		}
	}
}

func BenchmarkJoinWords(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		name := fmt.Sprintf("JoinWords-%d", size)
		s1 := randomPhrase(size)
		s2 := randomPhrase(size)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				JoinWords(s1, s2)
			}
		})
	}
}

// randomPhrase возвращает случайную фразу из n слов.
func randomPhrase(n int) string {
	const size = 3
	words := make([]string, n)
	for i := range words {
		words[i] = randomWord(size)
	}
	return strings.Join(words, " ")
}

var rnd = rand.New(rand.NewSource(0))

// randomWord возвращает случайное слово из n букв.
func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rnd.Intn(len(letters))]
	}
	return string(chars)
}