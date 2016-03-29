package main

import "fmt"

type IntSet struct {
	words []uint64
}

func main() {
	s := new(IntSet)
	result := s.AddAll(1, 2, 3)
	fmt.Printf("Add All  : %d\n", result)
}

func (s *IntSet) AddAll(vals ...int) int {
	var result int
	for _, v := range vals {
		result += v
	}
	return result
}
