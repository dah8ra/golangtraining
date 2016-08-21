package ex121

import "testing"

func Test(t *testing.T) {
	ar := [5]int{1, 2, 3, 4, 5}
	m := make(map[[5]int]string)
	m[ar] = "a"
	Display("m", m)
}
