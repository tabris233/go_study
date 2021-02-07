package main 

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
    var s = time.Now()
	var ret = int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	fmt.Println(time.Now().Sub(s), ret, "---")
	return ret
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
    var s = time.Now()
	var ret int = 0
	
	for i:=0; i<8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}

	fmt.Println(time.Now().Sub(s), ret, "+++")
	return ret
}

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
    var s = time.Now()
	var ret int = 0
	
	for ; x>0; x>>=1 {
		ret += int(x&1)
	}

	fmt.Println(time.Now().Sub(s), ret, "**")
	return ret
}

// PopCount4 returns the population count (number of set bits) of x.
func PopCount4(x uint64) int {
    var s = time.Now()
	var ret int = 0
	
	for ; x>0; x=(x&(x-1)) {
		ret ++
	}

	fmt.Println(time.Now().Sub(s), ret, "...")
	return ret
}

func main() {
	PopCount(2223234234234)
	PopCount2(2223234234234)
	PopCount3(2223234234234)
	PopCount4(2223234234234)
}