package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

/// NewReader

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

/// newreader end

type LimitR struct {
	R     io.Reader
	limit int64
}

func (l *LimitR) Read(p []byte) (int, error) {
	if l.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.limit {
		p = p[:l.limit]
	}
	n, err := l.R.Read(p)
	l.limit -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitR{r, n}
}

func main() {
	lr := LimitReader(NewReader("12345"), 2)

	s, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Println("err: ", s)
	}
	fmt.Println(string(s))

	s, err = ioutil.ReadAll(lr)
	if err != nil {
		fmt.Println("err: ", s)
	}
	fmt.Println(string(s))
}
