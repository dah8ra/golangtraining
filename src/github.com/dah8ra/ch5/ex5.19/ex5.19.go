package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		xx(path)
	}
}

// Connect to an url by turning off wi-fi or not
func xx(url string) {
	type expected struct{}
	defer func() {
		if p := recover(); p != nil {
			switch p {
			case nil:
				fmt.Println("no error\n")
			case expected{}:
				fmt.Printf("expected error\n")
			case 1:	
				fmt.Printf("expected error: %v\n", p)
			default:
				panic(p)
			}
		}
	}()
	fmt.Printf("url: %s\n", url)
	r, err := http.Get(url)
	// Change a condition below.	
	if err != nil {
		panic(1)
	} else {
		panic(expected{})
	}
	fmt.Println(r)
}
