package main

import (
	"bytes"
	"fmt"
)

func join2(sep string, strs ...string) string {
	ans := "" // TODO byte.Buffer 更加高效

	firstFlag := false
	for _, s := range strs {
		if firstFlag {
			ans += sep
		}
		firstFlag = true
		ans += s
	}

	return ans
}

func join3(sep string, strs ...string) string {
	var ans bytes.Buffer

	firstFlag := false
	for _, s := range strs {
		if firstFlag {
			ans.WriteString(sep)
		}
		firstFlag = true
		ans.WriteString(s)
	}

	return ans.String()
}

func main() {
	fmt.Println(join2(" ", "b", "c", "d", "e", "f"))
}
