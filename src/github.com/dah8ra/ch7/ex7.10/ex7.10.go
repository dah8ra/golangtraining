package main

import (
	"fmt"
	"sort"
)

func main() {
	if IsPalindrome(isPalindrome(row)) {
		fmt.Println("IsPalindrome returned \"true\"")
	} else {
		fmt.Println("IsPalindrome returned \"false\"")
	}
}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) {
			return false
		}
	}
	return true

}

type isPalindrome []*Target

type Target struct {
	s string
}

var row = []*Target{
	{"a"},
	{"b"},
	{"c"},
	{"b"},
	{"a"},
}

/*
var row = []*Target{
	{"a"},
	{"c"},
	{"a"},
	{"r"},
	{"a"},
	{"m"},
	{"a"},
	{"n"},
	{"a"},
	{"m"},
	{"a"},
	{"r"},
	{"a"},
	{"c"},
	{"a"},
}
*/
func (x isPalindrome) Len() int { return len(x) }
func (x isPalindrome) Less(i, j int) bool {
	fmt.Printf("%s %s\n", x[i], x[j])
	if x[i] == x[j] {
		return true
	}
	return false
}
func (x isPalindrome) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
