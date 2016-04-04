package main

import (
	"fmt"
	"os"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	s := os.Args[1:]
	for _, v := range s {
		fmt.Println(breadthFirst(v))
	}
}

func breadthFirst(target string) bool {
	i := 0
	for _, v := range prereqs {
		if len(v) > 0 {
			for _, r := range v {
				if target == r {
					return true
				}
			}
		}
		i++
	}
	return false
}
