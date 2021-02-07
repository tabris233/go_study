package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := make(map[string]int)
	in  := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		cnt[in.Text()]++
	}

	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n , line)
		}
	}
}