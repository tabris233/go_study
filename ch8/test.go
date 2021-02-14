package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		//
		go func() {
			time.Sleep(time.Second * 2)

			fmt.Println("-----------------------")
		}()
		fmt.Println("+++++++++++++++++++++++")
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("***********************")
	time.Sleep(time.Second * 10)
}
