package main

import (
	"fmt"
	"time"
)

var ch = make(chan bool)

func main() {

	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	// ch <- true
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// if c, ok := <-ch; ok { // chan 会阻塞goroutine。
			// 	return
			// }
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
