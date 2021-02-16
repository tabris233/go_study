package intset

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"testing"
)

type mapSet map[int]bool

func (m *mapSet) unionWith(t *mapSet) {
	for k, _ := range *t {
		(*m)[k] = true
	}
}

func mapToString(m mapSet) string {
	var buf bytes.Buffer
	buf.WriteString("{")
	intlist := make(sort.IntSlice, 0, len(m))
	for k := range m {
		intlist = append(intlist, k)
	}
	sort.Sort(intlist)
	flag := false
	for _, k := range intlist {
		if flag {
			buf.WriteString(" ")
		}
		flag = true
		buf.WriteString(strconv.Itoa(k))
	}
	buf.WriteString("}")
	return buf.String()
}

func initData(x, y *IntSet, xx, yy mapSet) {
	x.Add(1)
	xx[1] = true
	x.Add(144)
	xx[144] = true
	x.Add(9)
	xx[9] = true

	y.Add(9)
	yy[9] = true
	y.Add(42)
	yy[42] = true
}

func TestAdd(t *testing.T) {
	var x, y IntSet
	xx := make(mapSet)
	yy := make(mapSet)

	// func (s *IntSet) Add(x int)
	x.Add(1)
	xx[1] = true
	x.Add(144)
	xx[144] = true
	x.Add(9)
	xx[9] = true
	// func (s *IntSet) String() string
	// fmt.Println(x.String()) // "{1 9 144}"
	if x_str, xx_str := x.String(), mapToString(xx); x_str != xx_str {
		t.Errorf("x.String() [%s] != map(xx_str) [%s]", x_str, xx_str)
	}

	y.Add(9)
	yy[9] = true
	y.Add(42)
	yy[42] = true
	// fmt.Println(y.String()) // "{9 42}"
	if y_str, yy_str := y.String(), mapToString(yy); y_str != yy_str {
		t.Errorf("y.String() [%s] != map(yy_str) [%s]", y_str, yy_str)
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	xx := make(mapSet)
	yy := make(mapSet)
	initData(&x, &y, xx, yy)

	x.UnionWith(&y)
	xx.unionWith(&yy)
	// fmt.Println(x.String()) // "{1 9 42 144}"
	if x_str, xx_str := x.String(), mapToString(xx); x_str != xx_str {
		t.Errorf("x.String() [%s] != map(xx_str) [%s]", x_str, xx_str)
	}
}

func TestHas(t *testing.T) {
	var x, y IntSet
	xx := make(mapSet)
	yy := make(mapSet)
	initData(&x, &y, xx, yy)

	// func (s *IntSet) Has(x int) bool
	// fmt.Println(x.Has(9), x.Has(123)) // "true false"
	if x.Has(9) != xx[9] {
		t.Errorf("x.Has(9):[%t] != xx[9]:[%t]", x.Has(9), xx[9])
	}
}

func TestClear(t *testing.T) {
	var x, y IntSet
	xx := make(mapSet)
	yy := make(mapSet)
	initData(&x, &y, xx, yy)

	// func (s *IntSet) Clear()
	x.Clear()
	for k := range xx {
		delete(xx, k)
	}
	// fmt.Println(y.String()) // "{}"
	if x_str, xx_str := x.String(), mapToString(xx); x_str != xx_str {
		t.Errorf("x.String() [%s] != map(xx_str) [%s]", x_str, xx_str)
	}
}

func asfd() {
	var x, y IntSet
	xx := make(mapSet)
	yy := make(mapSet)
	initData(&x, &y, xx, yy)

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

	var a, b IntSet
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
