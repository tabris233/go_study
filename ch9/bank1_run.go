package main

import (
	"fmt"
	"time"

	"./bank"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
		if !bank.WithDraw(50) {
			fmt.Println("=", bank.Balance(), "less than", 50) // A2
		}
		fmt.Println("=", bank.Balance()) // A2
	}()

	// Bob:
	go func() {
		bank.Deposit(100)                // B
		fmt.Println("=", bank.Balance()) // A2
		if !bank.WithDraw(50) {
			fmt.Println("=", bank.Balance(), "less than", 50) // A2
		}
		fmt.Println("=", bank.Balance()) // A2
		if !bank.WithDraw(500) {
			fmt.Println("=", bank.Balance(), "less than", 500) // A2
		}
		fmt.Println("=", bank.Balance()) // A2
		if !bank.WithDraw(50) {
			fmt.Println("=", bank.Balance(), "less than", 50) // A2
		}
		fmt.Println("=", bank.Balance()) // A2
	}()

	time.Sleep(time.Second)
}
