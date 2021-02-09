package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer f(%d)\n", x)
	f(x - 1)
}

func main2() {
	defer printStack()
	defer fmt.Printf("------")
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func testPanic() (x int, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case 2:
			err = fmt.Errorf("2")
		default:
			panic(p)
		}
	}()

	x = 1

	if x == 2 {
		panic(2)
	} else {
		panic(new(int))
	}
}

func main() {
	fmt.Println(testPanic())
}
