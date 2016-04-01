package main

import "fmt"

//import "fmt"

type IntSet struct {
	words []uint64
}

func main() {
	s := IntSet{[]uint64{0, 1, 2, 3, 4, 5, 5, 7}}
	t := &IntSet{[]uint64{0, 1, 2, 3, 4, 3, 9}}
	fmt.Printf("Input      : %d\n", s)
	//	s.DifferenceWith(t)
	//	s.IntersectWith(t)
	s.SynmetricDifferenceWith(t)
	fmt.Printf("Result : %d\n", s)
}

func (s *IntSet) SynmetricDifferenceWith(t *IntSet) {
	res := make([]uint64, 0)
	for i := 0; i < len(s.words); i++ {
		if !memberWithSkip(s.words[i], s.words, i) {
			if !member(s.words[i], t.words) {
				res = append(res, s.words[i])
			}
		}
	}
	for i := 0; i < len(t.words); i++ {
		if !memberWithSkip(t.words[i], t.words, i) {
			if !member(t.words[i], s.words) {
				res = append(res, t.words[i])
			}
		}
	}
	s.words = res
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	res := make([]uint64, 0)
	for i := 0; i < len(t.words); i++ {
		if !member(s.words[i], t.words) {
			res = append(res, s.words[i])
		}
		//		fmt.Printf("res=%d\n", res)
	}
	s.words = res
}

func (s *IntSet) IntersectWith(t *IntSet) {
	res := make([]uint64, 0)
	for i := 0; i < len(t.words); i++ {
		if member(s.words[i], t.words) {
			res = append(res, s.words[i])
		}
		//		fmt.Printf("res=%d\n", res)
	}
	s.words = res
}

func member(n uint64, xs []uint64) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

func memberWithSkip(n uint64, xs []uint64, skipIndex int) bool {
	for i := 0; i < len(xs); i++ {
		if i == skipIndex {
			continue
		}
		if n == xs[i] {
			return true
		}
	}
	return false
}
