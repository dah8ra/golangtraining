package ex91_test

import (
	"fmt"
	"testing"

	"github.com/dah8ra/ch9/ex91"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		ex91.Deposit(200)
		fmt.Println("=", ex91.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		for {
			if ex91.Withdraw(100) {
				break
			}
		}
		done <- struct{}{}
	}()
	go func() {
		for {
			if ex91.Withdraw(10) {
				break
			}
		}
		done <- struct{}{}
	}()
	go func() {
		for {
			if ex91.Withdraw(20) {
				break
			}
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done

	if got, want := ex91.Balance(), 70; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
