package main

import (
	"fmt"

	"./intset"
)

func main() {
	var x, y intset.IntSet

	// func (s *IntSet) Add(x int)
	x.Add(1)
	x.Add(144)
	x.Add(9)
	// func (s *IntSet) String() string
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	// func (s *IntSet) UnionWith(t *IntSet)
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	// func (s *IntSet) Has(x int) bool
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// func (s *IntSet) Clear()
	y.Clear()
	fmt.Println(y.String()) // "{}"

	// func (s *IntSet) Copy() *IntSet
	y = *x.Copy()
	fmt.Println(y.String()) // "{1 9 42 144}"

	// func (s *IntSet) Len() int
	fmt.Println(x.Len()) // "4"
	fmt.Println(y.Len()) // "4"

	// func (s *IntSet) Remove(x int)
	x.Remove(9)
	fmt.Println(x.String()) // "{1 42 144}"
	x.Remove(9)
	fmt.Println(x.String()) // "{1 42 144}"

	// ----------------------------------------------------------------

	var a, b intset.IntSet
	// func (s *IntSet) IntersectWith(t *IntSet)
	a.Clear()
	b.Clear()
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	a.IntersectWith(&b)
	fmt.Println(a.String()) // "{2}"

	// func (s *IntSet) DifferenceWith(t *IntSet)
	a.Clear()
	b.Clear()
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	a.DifferenceWith(&b)
	fmt.Println(a.String()) // "{1}"

	// func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet
	a.Clear()
	b.Clear()
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	c := a.SymmetricDifference(&b)
	fmt.Println(c.String()) // "{1 3}"

	// func (s *IntSet) Elems()
	for i, elem := range c.Elems() {
		fmt.Println(i, elem)
	}
}
