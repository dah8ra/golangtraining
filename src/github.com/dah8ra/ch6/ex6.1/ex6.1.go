package main

import "fmt"

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	w := IntSet{[]uint64{0, 1, 2, 3, 4, 5, 2, 3, 4}}
	fmt.Printf("Before Remove: %d\n", w)
	fmt.Printf("Remove Point : %d\n", w.Remove(2))
	fmt.Printf("After Remove : %d\n", w)
	fmt.Printf("Before Copy  : %d\n", w)
	cp := w.Copy()
	fmt.Printf("After  Copy  : %d\n", cp)
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) int {
	if x > len(s.words) {
		return -1
	}
	t := make([]uint64, len(s.words)-1, 2*(len(s.words)-1))
	index := 0
	for i, v := range s.words {
		if i == x {
			continue
		} else {
			t[index] = v
		}
		index++
	}
	s.Clear()
	s.words = make([]uint64, len(t), 2*len(t))
	for i, v := range t {
		s.words[i] = v
	}
	return x
}

func (s *IntSet) Clear() {
	fmt.Printf("Before Clear -> %d\n", s.words)
	s.words = s.words[:0]
	fmt.Printf("After Clear  -> %d\n", s.words)
}

func (s *IntSet) Copy() *IntSet {
	t := make([]uint64, len(s.words), 2*(len(s.words)))
	temp := new(IntSet)
	temp.words = t

	for i, v := range s.words {
		temp.words[i] = v
	}
	return temp
}
