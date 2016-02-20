package main

import "fmt"

func main() {
	var src = []string{"a", "a", "b", "c", "c", "c", "d", "e"}
	fmt.Println(delete(src))
}

func deleteAndStuff(src []string) []string {
	var temp string
	result := []string{}
	for _, v := range src {
		if temp != v {
			result = append(result, v)
		}
		temp = v
	}
	return result
}

func delete(src []string) []string {
	var temp string
	for i, v := range src {
		if temp != v {
			src[i] = ""
		}
		temp = v
	}
	return src
}
