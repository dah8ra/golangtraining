package main

import (
	"fmt"

	"github.com/dah8ra/ch7/eval714"
)

func main() {
	expr, _ := eval714.Parse("1+2*3+pow(2,2)-4/8+sin(60)")
	fmt.Println(expr)
	env := eval714.Env{"x": 1, "y": 2}
	if err := expr.Check(nil) ; err != nil {
		fmt.Println("Error!")
		return
	}
	temp := expr.Calc(env)
	fmt.Printf("Minimum: %f\n", temp)
}
