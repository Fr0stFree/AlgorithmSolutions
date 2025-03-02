package main

import (
	"fmt"
	"io"
	"strings"
)

type AbyssWriter struct {
	total int
}

func (w *AbyssWriter) Write(b []byte) (int, error) {
	n, err := io.Discard.Write(b)
	if err != nil {
		return 0, err
	}
	w.total += n
	return n, nil

}

func (w *AbyssWriter) Total() int {
	return w.total
}

func NewAbyssWriter() *AbyssWriter {
	return &AbyssWriter{0}
}

func main() {
	r := strings.NewReader("go is awesome")
	w := NewAbyssWriter()
	written, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes\n", written)
	fmt.Println(written == int64(w.Total()))
}
