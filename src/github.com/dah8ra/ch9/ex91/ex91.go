package ex91

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	before := Balance()
	withdraws <- amount
	after := Balance()
	fmt.Println(after)
	if after == (before - amount) {
		return true	
	} else {
		return false
	}
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-withdraws:
			balance -= amount
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
