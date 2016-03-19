package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string]map[string]string{
	"algorithms": {"data structures": ""},
	"calculus":   {"linear algebra": ""},

	"compilers": {
		"data structures":       "",
		"formal languages":      "",
		"computer organization": "",
	},

	"data structures":  {"discrete math": ""},
	"databases":        {"data structures": ""},
	"discrete math":    {"intro to programming": ""},
	"formal languages": {"discrete math": ""},
	"networks":         {"operating systems": ""},
	"operating systems": {
		"data structures":       "",
		"computer organization": ""},
	"programming languages": {
		"data structures":       "",
		"computer organization": ""},
}

func main() {
	i := 0
	for key, _ := range topoSort(prereqs) {
		if len(key) != 0 {
			fmt.Printf("%d:\t%s\n", i+1, key)
			i++
		}
	}
}

func topoSort(m map[string]map[string]string) map[string]string {
	var order map[string]string = make(map[string]string)
	seen := make(map[string]bool)
	var visitAll func(items map[string]map[string]string)

	visitAll = func(items map[string]map[string]string) {
		for item, imap := range items {
			if !seen[item] {
				seen[item] = true
				var nitems map[string]map[string]string = make(map[string]map[string]string)
				var nmap map[string]string = make(map[string]string)
				for k, v := range imap {
					nmap[v] = ""
					nitems[k] = make(map[string]string)
					nitems[k] = nmap
				}
				visitAll(nitems)
				order[item] = ""
			}
		}
	}

	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var target map[string]map[string]string = make(map[string]map[string]string)
	for _, key := range keys {
		target[key] = make(map[string]string)
		target[key] = m[key]
	}
	visitAll(target)
	return order
}
