package eval713

import "fmt"

var index = 0

func (v Var) String(env Env) {
	tab := ""
	for i := 0; i < index; i++ {
		tab += " "
	}
	fmt.Printf("%s", tab)
	fmt.Println(v)
}

func (l literal) String(env Env) {
	tab := ""
	for i := 0; i < index; i++ {
		tab += " "
	}
	fmt.Printf("%s", tab)
	fmt.Println(l)
}

func (u unary) String(env Env) {
	fmt.Println("unary")
	tab := ""
	for i := 0; i < index; i++ {
		tab += " "
	}
	switch u.op {
	case '+':
		fmt.Printf("%s+%s\n", tab,u.x.Eval(env))
	case '-':
		fmt.Printf("%s-%f\n", tab,u.x.Eval(env))
	}
}

func (b binary) String(env Env) {
	tab := ""
	for i := 0; i < index; i++ {
		tab += " "
	}
	index++
	switch b.op {
	case '+':
		fmt.Printf("%s+\n", tab)
		b.x.String(env)
		b.y.String(env)
	case '-':
		fmt.Printf("%s-\n", tab)
		b.x.String(env)
		b.y.String(env)
	case '*':
		fmt.Printf("%s*\n", tab)
		b.x.String(env)
		b.y.String(env)
	case '/':
		fmt.Printf("%s/\n", tab)
		b.x.String(env)
		b.y.String(env)
	}
	index--
	//	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String(env Env) {
	tab := ""
	for i := 0; i < index; i++ {
		tab += " "
	}
	switch c.fn {
	case "pow":
		fmt.Printf("%spow(%f,%f)\n", tab,c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		fmt.Printf("%ssin(%f)\n", tab,c.args[0].Eval(env))
	case "sqrt":
		fmt.Printf("%ssqrt(%f)\n", tab,c.args[0].Eval(env))
	}
}
