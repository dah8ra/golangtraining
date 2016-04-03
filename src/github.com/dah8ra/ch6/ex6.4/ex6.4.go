package main

import "fmt"

type IntSet struct {
	words []uint64
}

func main() {
	s := IntSet{[]uint64{0, 1, 2, 3, 4, 5, 5, 7}}
	fmt.Printf("Input      : %d\n", s)
	fmt.Printf("Result : %d\n", s.Elems())
}

func (s *IntSet) Elems() []uint64 {
	res := make([]uint64, 0)
	for i := 0; i < len(s.words); i++ {
		res = append(res, s.words[i])
	}
	return res
}
