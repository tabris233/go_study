package main

import "fmt"

func rotate(a []int, r int) []int {
	la := len(a)
	r = r % la

	ret := make([]int, la)
	for i, v := range a {
		ret[(i+r)%la] = v
	}

	return ret
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)

	fmt.Println(rotate(a, 1))
	fmt.Println(rotate(a, 2))
	fmt.Println(rotate(a, 3))
}
