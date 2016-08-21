package ex122

import (
	"os"
	"reflect"
	"testing"
)

// This test ensures that the program terminates without crashing.
func Test(t *testing.T) {
	// Even metarecursion!  (YMMV)
	Display("rV", reflect.ValueOf(os.Stderr))
	// Output:
	// Display rV (reflect.Value):
	// (*rV.typ).size = 8
	// (*rV.typ).ptrdata = 8
	// (*rV.typ).hash = 871609668
	// (*rV.typ)._ = 0
	// ...

	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if false {
		Display("c", c)
		// Output:
		// Display c (display.Cycle):
		// c.Value = 42
		// (*c.Tail).Value = 42
		// (*(*c.Tail).Tail).Value = 42
		// ...ad infinitum...
	}

}
