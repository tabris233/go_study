// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // receive withdraw
var isLess = make(chan bool)   // is
func Deposit(amount int)       { deposits <- amount }
func Balance() int             { return <-balances }
func WithDraw(amount int) bool {
	withdraws <- amount
	return <-isLess
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if amount <= balance {
				balance -= amount
				isLess <- true
			} else {
				isLess <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
