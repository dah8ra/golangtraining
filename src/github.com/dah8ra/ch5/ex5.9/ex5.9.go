package main

import (
	"flag"
	"fmt"
	"strings"
)

var target = "I$fooneed$foomore$footime"
var n = flag.Bool("n", false, "omit trailing newline")
var w = flag.String("s", "foo", "search word")

func main() {
	flag.Parse()
	fmt.Printf("Search query: %s\n", *w)
	fmt.Printf("Target: %s\n", target)
	ret := expand(*w, x)
	fmt.Printf("Result: %s\n", ret)
}

func expand(s string, f func(string) string) string {
	temp := f(s)
	sep := "$" + s
	sarray := strings.Split(target, sep)
	//	fmt.Printf("%s\n", sarray)
	return interpolateCombine(sarray, temp)
	//	return expand(temp, f)
}

func interpolateCombine(sarray []string, s string) string {
	//	fmt.Printf("Base word: %s\n", sarray)
	//	fmt.Println("Interpolate word: " + s)
	var ret string
	for i := 0; i < len(sarray)*2; i++ {
		//		fmt.Printf("Index: %d %d\n", i, len(sarray)*2)
		if i == len(sarray)*2-2 {
			return ret
		} else if i%2 == 0 {
			ret += sarray[i/2]
		} else if i%2 == 1 {
			ret += s
		}
	}
	return ret
}

func firstSkipCombine(sarray []string) string {
	i := 0
	var ret string
	for _, str := range sarray {
		if i != 0 {
			ret += str
		}
		i++
	}
	return ret
}

func x(s string) string {
	temp := "$" + s
	sarray := strings.Split(target, temp)
	//	fmt.Printf("Split result: %s\n", sarray)
	//	if len(sarray) == 1 {
	//		fmt.Printf("ret no length, %s\n", target)
	//		return sarray[0]
	//	}
	//	ret := firstSkipCombine(sarray)
	//	fmt.Printf("Combined %s\n", string(ret))
	//	target = string(ret)
	length := len(sarray)
	separray := sarray[length-1 : length]
	return separray[0]
}
