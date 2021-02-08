package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func test() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	ages = make(map[string]int)

	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	ages["a"] = 1

	age, ok := ages["a"]
	fmt.Println(age, ok)
	age, ok = ages["b"]
	fmt.Println(age, ok)
}

func main() {
    //counts := make(map[string]map[rune]int)    // counts of Unicode characters
    var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
    invalid := 0                    // count of invalid UTF-8 characters

    letters := make(map[rune]int)
    numbers := make(map[rune]int)
    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune() // returns rune, nbytes, error
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        /*
        练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
        */
        //判断是字母
        if unicode.IsLetter(r){
                letters[r]++
        }
        //判断是数字
        if unicode.IsNumber(r){
                numbers[r]++
        }
        //counts[r]++
        utflen[n]++
    }
    fmt.Printf("rune\tcount\n")
    for c, n := range letters {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Printf("rune(number)\tcount\n")
    for c, n := range numbers {
        fmt.Printf("%q\t%d\n", c, n)
    }

    fmt.Print("\nlen\tcount\n")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}