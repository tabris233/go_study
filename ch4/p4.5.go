package main

import "fmt"

func f(strs []string) []string {
	p := 1

	for i, s := range strs {
		if i > 0 {
			if s != strs[i-1] {
				strs[p] = s
				p++
			}
		}
	}

	strs = strs[:p]

	return strs
}

func main() {
	strs := []string{"foo", "bar", "baz", "baz", "a", "b", "c", "c", "c", "d"}

	fmt.Println(strs)
	// f(strs)
	// fmt.Println(strs)
	fmt.Println(f(strs))
}
