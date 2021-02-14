package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type S struct {
	R       string
	current int
}

func (s *S) Read(p []byte) (int, error) {
	if s.current >= len(s.R) {
		return 0, io.EOF
	}
	l := copy(p, s.R[s.current:])
	s.current = l
	fmt.Println("---", l)
	return l, nil
}

func NewReader(s string) io.Reader {
	return &S{s, 0}
}

func main() {
	r := NewReader("1234")
	s, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(s)
	}
	fmt.Println(string(s))
}
