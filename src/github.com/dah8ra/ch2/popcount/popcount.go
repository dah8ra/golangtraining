package popcount

import (
	"fmt"
)

var pc [256]byte

func init() {
	fmt.Println("Call init()")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Print(byte(i&1))
		//fmt.Print(i/2)
		fmt.Print(pc[i])
	}
}

func Test(x uint64) int {
	return int(pc[byte(x>>(0*8))])
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
