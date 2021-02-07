package main

import (
	"fmt"
	"unicode"
)

func f(a []byte) []byte {
	for i := 0; i < len(a)-1; i++ {
		if unicode.IsSpace(rune(a[i])) && unicode.IsSpace(rune(a[i+1])) {
			copy(a[i:], a[i+1:])
			a = a[:len(a)-1]
			i--
		}
	}

	return a
}

func main() {
	a := []byte("aasdfas   asdf asf s sf sa fasdf    sfdas")

	fmt.Println(string(a))
	fmt.Println(string(f(a)))
}
