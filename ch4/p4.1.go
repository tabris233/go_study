package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a, b := 5, 6
	c := a ^ b
	fmt.Printf("%b &^ %b = %b", a, b, c)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	cnt := 0
	for i, _ := range c1 {
		x := c1[i] ^ c2[i]
		for ; x > 0; x = x & (x - 1) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
