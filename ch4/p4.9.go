package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	cnts := make(map[string]int)
	for in.Scan() {
		cnts[in.Text()]++
	}

	for k, v := range cnts {
		fmt.Println("map[", k, "] == ", v)
	}

}