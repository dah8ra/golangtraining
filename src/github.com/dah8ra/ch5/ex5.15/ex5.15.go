package main

import "fmt"

func main() {

	fmt.Println(max())
	fmt.Println(max(3))
	fmt.Println(max(1, 2, 3, 4))

	fmt.Println(min())
	fmt.Println(min(3))
	fmt.Println(min(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...))
	fmt.Println(min(values...))
}

func max(vals ...int) int {
	temp := 0
	for _, val := range vals {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func min(vals ...int) int {
	temp := 0
	isInit := false
	for _, val := range vals {
		if !isInit {
			temp = val
			isInit = true
		} else if val < temp {
			temp = val
		}
	}
	return temp
}
