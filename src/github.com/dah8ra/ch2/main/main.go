package main

import (
	"fmt"
	"github.com/dah8ra/ch2/tempconv"
)

func main() {
	//var c tempconv.Celsius = 100.0 	
	fmt.Printf("Brrrr %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
