package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{"a", "b", "c"}
	fmt.Println(xx(ss, "d", "e"))
}

func xx(a []string, seps ...string) string {
	var temp string
	for _, sep := range seps {
		temp = strings.Join(a, sep)
		b := []byte(temp)
		i := 0
		a = make([]string, len(temp))
		for _, t := range b {
			a[i] = string(t)
			i++
		}
		fmt.Printf("@@@ %v\n", temp)
	}
	return temp
}
