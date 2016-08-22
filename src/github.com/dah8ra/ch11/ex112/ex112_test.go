package ex112

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1, "{1}"},
		{144, "{1 144}"},
		{9, "{1 9 144}"},
	}
	var x IntSet
	fmt.Println("===========")
	for _, test := range tests {
		x.Add(test.input)
		if string(x.String()) != test.want {
			t.Errorf("Add(%d) = %v", test.input, x.String())
		}
	}
	fmt.Println("===========")
}

func TestIntSet(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	m := make(map[int]uint)
	slice := []int{1, 144, 9}
	for i := 0; i < len(slice); i++ {
		bit := uint(slice[i] % 64)
		word := slice[i] / 64
		t := m[word]
		if t != 0 {
			//m[word] = t + bit
			m[word] |= 1 << bit
		} else {
			m[word] |= 1 << bit
		}
		fmt.Println("m: ", m)
	}

	for i := 0; i < len(m); i++ {
		//fmt.Println("### ", x.words[i])
		//fmt.Println("@@@ ", m[i])
		//if x.words[i] == m[i] {
		//			t.Error("Error!!!")
		//	}
	}
}

/*
func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
*/
