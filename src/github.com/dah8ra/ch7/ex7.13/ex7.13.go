package main

import (
	"fmt"

	"github.com/dah8ra/ch7/eval"
)

func main() {
	expr, _ := eval.Parse("1+2*pow(x,y)-3/4+sin(7)")
	fmt.Println(expr)
	env := eval.Env{"x": 1, "y": 2}
	expr.String(env)
}
