package main

import (
	"github.com/dah8ra/ch7/treesort"
//	"math/rand"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		//		data[i] = rand.Int() % 50
		data[i] = i
	}
	root := treesort.Sort(data)
	root.String()
}
