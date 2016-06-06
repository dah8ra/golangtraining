package main

import (
	"fmt"
	"github.com/dah8ra/ch7/treesort"
	"math/rand"
)

func main() {
	data := make([]int, 8)
	for i := range data {
		data[i] = rand.Int() % 8
	}
	//	data[0] = 4
	//	data[1] = 3
	//	data[2] = 2
	//	data[3] = 8
	//	data[4] = 6
	//	data[5] = 9
	fmt.Printf("INPUT: %v\n", data)
	root := treesort.Sort(data)
	root.String()
}
