package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Println(fmt.Sprintf("Index: %d, Value: "+v, i))
	}
}
