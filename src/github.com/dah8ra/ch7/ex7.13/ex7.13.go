package main

import (
	"fmt"

	"github.com/dah8ra/ch7/eval713"
)

func main() {
	expr, _ := eval713.Parse("1+2*3-4/8")
	fmt.Println(expr)
	env := eval713.Env{"x": 1, "y": 2}
	expr.String(env)	
}
