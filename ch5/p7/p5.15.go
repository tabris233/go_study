package main

import (
	"fmt"
	"math"
)

func max(args ...int) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("no args")
	}
	mx := math.MinInt64
	for _, v := range args {
		if v > mx {
			mx = v
		}
	}
	return mx, nil
}

func min(args ...int) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("no args")
	}
	mn := math.MaxInt64
	for _, v := range args {
		if v < mn {
			mn = v
		}
	}
	return mn, nil
}

func main() {
	fmt.Println(max())
	fmt.Println(max(4))
	fmt.Println(max(4, 1, 3, 2))

	fmt.Println(min())
	fmt.Println(min(4))
	fmt.Println(min(4, 1, 3, 2))
}
