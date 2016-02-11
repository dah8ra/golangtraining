package main

import (
	"time"
	"fmt"
	//"strconv"
)

var pc [256]byte
var pc1 [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		pc1[i] = pc1[i/2] + byte(i&1)
	}
}

func main() {
	start1 := time.Now()
	fmt.Printf("Result        : %d\n", popcount(10))
	//pcSec := time.Since(start1).Seconds()
	pcSec := time.Since(start1).Nanoseconds()
	start2 := time.Now()
	fmt.Printf("Result by loop: %d\n", popcountByLoop(10))
	//pcByLoopSec := time.Since(start2).Seconds()
	pcByLoopSec := time.Since(start2).Nanoseconds()
	elapsed := pcSec - pcByLoopSec
	//fmt.Printf(strconv.FormatFloat(elapsed, 'G', 4, 64))
	fmt.Println()
	fmt.Printf("%dnano sec elapsed.", elapsed)
}

func popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))]) 
}

func popcountByLoop(x uint64) int {
	var pc int
	var i uint
	for i=0 ; i<8 ; i++ {
		pc += int(pc1[byte(x>>(i*8))])
	}
	return pc
}
