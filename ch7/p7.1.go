package main

import (
	"fmt"
)

// WordCounter count
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	// fmt.Println(p)
	l := 0
	p = append(p, byte(' '))
	for i, c := range p {
		if c == byte(' ') && i != 0 && p[i-1] != ' ' {
			l++
		}
	}
	*c += WordCounter(l) // convert int to WordCounter
	return l, nil
}

// LineCounter count
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	// fmt.Println(p)
	l := 0
	for _, c := range p {
		if c == byte('\n') {
			l++
		}
	}
	*c += LineCounter(l) // convert int to LineCounter
	return l, nil
}

func main() {
	var wc WordCounter
	fmt.Fprintf(&wc, "hasdf asf asdf asf sdf sdf %s", "asfasf sfasf 世界")
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, "hasdf asf asdf \n asf sdf sdf\n %s", "asfasf sfas\nf 世界")
	fmt.Println(lc)
}
