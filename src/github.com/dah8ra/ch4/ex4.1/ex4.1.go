package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := 0
	for i, v1 := range c1 {
		v2 := c2[i]
		if v1 != v2 {
			var v byte = v1 & v2
			fmt.Printf("%b %b -> %b\n", v1, v2, v)
			count += PopCountByClearing(v)
		}
	}
	fmt.Printf("TOTAL COUNT: %d\n", count)
}

func PopCountByClearing(x byte) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
