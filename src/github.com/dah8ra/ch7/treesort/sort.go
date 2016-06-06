package treesort

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() {
	show(t, false, -1)
}

func show(t *tree, init bool, indent int) {
	if t == nil {
		return
	}
	tab := ""
	for i := 0; i < indent; i++ {
		tab = tab + " "
	}
	if !init {
		fmt.Printf("%d\n", t.value)
	} else {
		fmt.Printf("%s├%d\n", tab, t.value)
	}
	if t.left != nil {
		show(t.left, true, indent+1)
	}
	if t.right != nil {
		show(t.right, true, indent+1)
	}
}
