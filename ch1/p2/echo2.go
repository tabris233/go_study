package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var st1 = time.Now()
	var s, sep string
	sep = "\n"
	for i, arg := range os.Args[0:] {
		s += strconv.Itoa(i) + " " + arg + sep
	}
	var e1 = time.Now()
	fmt.Println(s, e1.Sub(st1))

	var st2 = time.Now()
	var ss = strings.Join(os.Args, "\n")
	var e2 = time.Now()
	fmt.Println(ss, "\n", e2.Sub(st2))
}