package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////
// 235,904 communication per seconds
// OS : Win7 64bit
// CPU: Core i7 3.4GHz
// Mem: 4GB
/////////////////////////////////////
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			tick := time.Tick(1 * time.Second)
			select {
			case <-tick:
				fmt.Println("===========================")
			}
		}
	}()

	go func() {
		ch1 <- "pong"
		for {
			select {
			case msg := <-ch2:
				fmt.Println(msg)
				ch1 <- "pong"
			}
		}
	}()

	for msg := range ch1 {
		fmt.Println(msg)
		ch2 <- "ping"
	}

}
