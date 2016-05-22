package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dah8ra/ch7/eval715"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fn := "sqrt(x)"
	fmt.Printf("Expr: %s\n", fn)
	fmt.Println("Please input some values about x and y.")
	fmt.Println("For example => 1,2")
	fmt.Print("=> ")
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	var env map[eval715.Var]float64
	var expr eval715.Expr
	vars := make(map[eval715.Var]bool)
	if strings.Contains(s, ",") {
		val := strings.Split(s, ",")
		expr, _ = eval715.Parse(fn)
		x, _ := strconv.ParseFloat(val[0], 64)
		y, _ := strconv.ParseFloat(val[1], 64)
		env = eval715.Env{"x": x, "y": y}
	} else {
		expr, _ = eval715.Parse(fn)
		x, _ := strconv.ParseFloat(s, 64)
		env = eval715.Env{"x": x}
	}
	if err := expr.Check(vars); err != nil {
		return
	}
	answer := expr.Eval(env)
	fmt.Printf("Expr: %s\tAnswer: %f\n", fn, answer)
}
