// Interface Values With Nil Underlying Values:

// If the concrete values inside the interface itself is nil, the method will be called with a nil receiver.


// In some languages this would trigger a null pointer exception.

// In Go it is common to write methods that gracefully handl being called with a nil receiver.

// For example, with the method M in this program.


// Note that an interface value that holds a nil concrete value is itself non-nil.

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
