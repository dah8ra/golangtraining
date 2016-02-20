package main

import "fmt"

func main() {
	var src = []string{"a", "a", "b", "c","c","c","d","e"}
	var temp string
	result := []string{}
	for _, v := range src {
		if temp != v {
			result = append(result, v)
		}
		temp = v
	}
	fmt.Println(result)
}
