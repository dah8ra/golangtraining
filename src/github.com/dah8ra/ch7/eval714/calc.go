package eval714

import (
	"fmt"
	"math"
)

//!+Calc1

func (v Var) Calc(env Env) float64 {
	return env[v]
}

func (l literal) Calc(_ Env) float64 {
	return float64(l)
}

//!-Calc1

//!+Calc2

func (u unary) Calc(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Calc(env)
	case '-':
		return -u.x.Calc(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

var min float64 = 1.797693134862315708145274237317043567981e+308

func (b binary) Calc(env Env) float64 {
	switch b.op {
	case '+':
		temp := b.x.Calc(env) + b.y.Calc(env)
		if temp < min {
			min = temp
		}
		return min
	case '-':
		temp := b.x.Calc(env) - b.y.Calc(env)
		if temp < min {
			min = temp
		}
		return min
	case '*':
		temp := b.x.Calc(env) * b.y.Calc(env)
		if temp < min {
			min = temp
		}
		return min
	case '/':
		temp := b.x.Calc(env) / b.y.Calc(env)
		if temp < min {
			min = temp
		}
		return min
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Calc(env Env) float64 {
	switch c.fn {
	case "pow":
		temp := math.Pow(c.args[0].Calc(env), c.args[1].Calc(env))
		fmt.Printf("Fn: %v\tMin: %f\tVal: %f\n", c.fn, min, temp)
		if temp < min {
			min = temp
		}
		return min
	case "sin":
		temp := math.Sin(c.args[0].Calc(env))
		fmt.Printf("Fn: %v\tMin: %f\tVal: %f\n", c.fn, min, temp)
		if temp < min {
			min = temp
		}
		return min
	case "sqrt":
		temp := math.Sqrt(c.args[0].Calc(env))
		fmt.Printf("Fn: %v\tMin: %f\tVal: %f\n", c.fn, min, temp)
		if temp < min {
			min = temp
		}
		return min
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (c calc) Calc(env Env) float64 {
	switch c.op {
	case '+':
		temp := c.x.Calc(env) + c.y.Calc(env)
		fmt.Printf("Op: %v\tMin: %f\tVal: %f\r\n", c.op, min, temp)
		if temp < min {
			min = temp

		}
		//		fmt.Printf("After  Op: %v\tMin: %f\r\n", c.op, min)
		return min
	case '-':
		temp := c.x.Calc(env) - c.y.Calc(env)
		fmt.Printf("Op: %v\tMin: %f\tVal: %f\r\n", c.op, min, temp)
		if temp < min {
			min = temp
		}
		//		fmt.Printf("After  Op: %v\tMin: %f\r\n", c.op, min)
		return min
	case '*':
		temp := c.x.Calc(env) * c.y.Calc(env)
		fmt.Printf("Op: %v\tMin: %f\tVal: %f\r\n", c.op, min, temp)
		if temp < min {
			min = temp
		}
		//		fmt.Printf("After  Op: %v\tMin: %f\r\n", c.op, min)
		return min
	case '/':
		temp := c.x.Calc(env) / c.y.Calc(env)
		fmt.Printf("Op: %v\tMin: %f\tVal: %f\r\n", c.op, min, temp)
		if temp < min {
			min = temp
		}
		//		fmt.Printf("After  Op: %v\tMin: %f\r\n", "+", min)
		return min
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", c.op))
}

//!-Calc2
