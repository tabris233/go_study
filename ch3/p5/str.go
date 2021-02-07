package main

import (
	"bytes"
	"fmt"
	_ "strconv"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if s[0] == '-' || s[0] == '+' {
		buf.WriteString(s[0:1])
		s = s[1:]
	}

	arr := strings.Split(s, ".")
	s = arr[0]

	// pre of '.'
	preFlag := false

	ls := len(s)

	if ls%3 > 0 {
		preFlag = true
		buf.WriteString(s[:ls%3])
	}

	for i := ls % 3; i < ls; i += 3 {
		if preFlag {
			buf.WriteString(",")
		}
		preFlag = true
		buf.WriteString(s[i : i+3])
	}

	// suf of '.'
	if len(arr) > 1 {
		buf.WriteString(".")
		buf.WriteString(arr[1])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234.5678"))
	fmt.Println(comma("123412.567891"))
	fmt.Println(comma("-1234.5678"))
	fmt.Println(comma("+1234.5678"))
}
