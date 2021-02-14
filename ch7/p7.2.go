package main

import (
	"fmt"
	"io"
	"os"
)

/// type
type CountWriter struct {
	cw io.Writer
	C  int64
}

func (c *CountWriter) Write(p []byte) (n int, err error) {
	n, err = c.cw.Write(p)
	c.C = int64(n)
	return
}

/// type end

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CountWriter{cw: w}
	return c, &c.C // CountWriter 实现了Write函数， 也可以是io.Writer 接口类型了。
}

func main() {
	rw, ilen := CountingWriter(os.Stdout)
	rw.Write([]byte("hello, world!"))
	fmt.Println(" ", *ilen)
}
