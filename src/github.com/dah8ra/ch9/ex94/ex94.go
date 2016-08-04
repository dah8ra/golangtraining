package main

import (
	"fmt"
	"time"
)

var ch chan int = make(chan int)
var index int = 0

///////////////////////////////////////////////////
// Created 205,842 goroutines in 205.842 seconds.
// OS : Win7 64bit
// CPU: Core i7 3.4GHz
// Mem: 4GB
///////////////////////////////////////////////////
func main() {
	go receiver(ch)
	for {
		go thread(index)
		time.Sleep(1 * time.Millisecond)
		index++
	}
}

func thread(index int) {
	ch <- index
}

func receiver(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		}
	}
}
