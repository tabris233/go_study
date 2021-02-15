package main

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var InitPc sync.Once

func Init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	InitPc.Do(Init)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var wg sync.WaitGroup
	for i := 2223234234234; i < 2223234234234+20; i++ {
		wg.Add(1)
		go func(x uint64) {
			wg.Done()
			fmt.Println(PopCount(x))
		}(uint64(i))
	}

	wg.Wait()
}
