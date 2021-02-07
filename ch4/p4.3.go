package main

import "fmt"

func reverse(sPtr *[4]int) {
	for i, j := 0, len(*sPtr)-1; i < j; i, j = i+1, j-1 {
		sPtr[i], sPtr[j] = sPtr[j], sPtr[i]
	}
}

func main() {
	s := [...]int{1, 2, 3, 4}
	// s := [...]int{1, 2, 3, 4, 5, 6} [6]int not [4]int
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}
