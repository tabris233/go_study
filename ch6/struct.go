package main

import (
	"fmt"
	"math"
)

// Point struct
type Point struct {
	X, Y float64
}

// Distance Point method
func (p Point) Distance(q Point) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// Test struct
type Test struct {
	Point
	Z int
}

// Test2 struct
type Test2 struct {
	*Point
	Z int
}

// Test3 struct
type Test3 struct {
	p Point
	Z int
}

func test1() {
	t1 := Test{Point{1, 1}, 2}
	t2 := Test{Point{4, 5}, 3}
	fmt.Println(t1.Distance(t2.Point))
}

func test2() {
	t1 := Test2{&Point{1, 1}, 2}
	t2 := Test2{&Point{4, 5}, 3}
	fmt.Println(t1.Distance(*t2.Point))

	fmt.Printf("%#v", t1.Point)
}

func test3() {
	t1 := Test3{Point{1, 1}, 2}
	t2 := Test3{Point{4, 5}, 3}
	fmt.Println(t1.Distance(t2.p))
}

func main() {
	// test1()
	// test2()
	test3()
}
