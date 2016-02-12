package main

import (
	"fmt"
	"time"
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
	var input uint64 = 12
	elapsed := time.Now()
	fmt.Printf("Result        : %d\n", popcount(input))
	fmt.Printf("%dnano sec elapsed.", time.Since(elapsed).Nanoseconds)
}

func popcount(x uint64) int {
	y := (x&(x-1) | x&(-x))
	fmt.Println(x & (x - 1))
	fmt.Println(x & (-x))

	return int(pc[byte(y>>(0*8))] +
		pc[byte(y>>(1*8))] +
		pc[byte(y>>(2*8))] +
		pc[byte(y>>(3*8))] +
		pc[byte(y>>(4*8))] +
		pc[byte(y>>(5*8))] +
		pc[byte(y>>(6*8))] +
		pc[byte(y>>(7*8))])
}
