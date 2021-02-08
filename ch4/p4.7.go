package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(a []byte) []byte {
	for l := len(a); l > 0; {
		r, size := utf8.DecodeRune(a[0:])
		// fmt.Println(a)
		copy(a[0:l], a[0+size:l])
		// fmt.Println(a)
		copy(a[l-size:l], []byte(string(r)))
		// fmt.Println(a)
		// fmt.Println("-------")
		l -= size
	}

	return a
}

func main() {
	a := []byte("hello, 世界!")

	fmt.Println(string(a))
	fmt.Println(string(reverse(a)))
}
