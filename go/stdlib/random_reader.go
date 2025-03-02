package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

func RandomReader(max int) io.Reader {
	return &randomReader{max: max}
}

type randomReader struct {
	max   int
	index int
}

func (r *randomReader) Read(p []byte) (n int, err error) {
	if r.index >= r.max {
		return 0, io.EOF
	}

	remaining := r.max - r.index
	if len(p) > remaining {
		p = p[:remaining]
	}

	n, err = rand.Read(p)
	if err != nil {
		return 0, err
	}

	r.index += n

	return n, nil
}

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}
