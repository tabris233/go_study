package main

import (
	"fmt"
	"sort"
)

type st []string

func (s st) Len() int           { return len(s) }
func (s st) Less(i, j int) bool { return s[i] < s[j] }
func (s st) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	s1 := st{"a", "b", "c", "d", "c", "b", "a"}
	s2 := st{"a", "b", "e", "d", "c", "b", "a"}
	fmt.Println(s1, IsPalindrome(s1))
	fmt.Println(s2, IsPalindrome(s2))
}
