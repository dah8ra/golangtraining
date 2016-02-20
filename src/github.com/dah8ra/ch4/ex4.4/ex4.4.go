package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 3)
	fmt.Println(a)
}

func rotate(p []int, rotate int) {
	reverse(p[:rotate])
	reverse(p[rotate:])
	reverse(p[:])
}

func reverse(p []int) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}
