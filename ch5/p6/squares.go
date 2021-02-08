package main

import "fmt"

func squares() func() int {
	var x int
	x++
	return func() int {
		return x * x
	}
}

func squares2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	/********************************
	因为x是在squares中++后变成了1，后面匿名函数调用时x都是1
	*/
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("--------------------------------")

	/********************************
	因为x是在匿名函数中++， 每次调用匿名函数时x都加1， 且这时候还是f = squares2()第一次调用的x
	*/
	f = squares2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("--------------------------------")

	/********************************
	因为x是在匿名函数中++， 每次调用匿名函数时x都加1， 且这时候还是f = squares2()时的squeare2()的x
	*/
	f = squares2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("--------------------------------")

	/********************************
	因为x是在匿名函数中++， 每次调用匿名函数时x都加1， 且这时候还是f = squares2()时的squeare2()的x
	squares2() 返回的匿名函数没有被使用，
	f 还是之前那次 squares2() 返回的匿名函数
	*/
	squares2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
