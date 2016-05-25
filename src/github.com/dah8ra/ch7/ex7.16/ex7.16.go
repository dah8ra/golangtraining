package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dah8ra/ch7/eval716"
)

// Sample query
// http://localhost:8000/calc?expr=pow(x,y)&x=2&y=2

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		calc(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calc(w http.ResponseWriter, req *http.Request) {
	expr_query := req.URL.Query().Get("expr")
	x_query := req.URL.Query().Get("x")
	y_query := req.URL.Query().Get("y")

	fmt.Printf("%s\n%s\n%s\n", expr_query, x_query, y_query)

	var env map[eval716.Var]float64
	var expr eval716.Expr
	vars := make(map[eval716.Var]bool)
	if x_query != "" && y_query != "" {
		expr, _ = eval716.Parse(expr_query)
		x, _ := strconv.ParseFloat(x_query, 64)
		y, _ := strconv.ParseFloat(y_query, 64)
		env = eval716.Env{"x": x, "y": y}
	} else {
		expr, _ = eval716.Parse(expr_query)
		x, _ := strconv.ParseFloat(x_query, 64)
		env = eval716.Env{"x": x}
	}

	if err := expr.Check(vars); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	answer := expr.Eval(env)
	fmt.Fprintf(w, "Expr: %s\tAnswer: %f\n", expr, answer)
}
