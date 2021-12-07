// Pointers
// Go has pointers. A pointer holds the memory address of a value.

// The type *T is a pointer to a value. Its zero value is nil.

//	var p *int

// The & operator generates a pointer to its operand.

//	i := 42
//	p = &i

// The * operator denotes the pointer's underlying value.

//	fmt.Println(*p) // Read i through the pointer p.
//	*p = 21		// Set i through the pointer p.

// This is known as "dereferencing" or "indirecting".
// Unlike C, Go has no pointer arithmetic.

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i		// Point to i.
	fmt.Println(*p)	// Read i through the pointer.
	*p = 21		// Set i through the pointer.
	fmt.Println(i)	// See the new value of i.

	p = &j		// Point to j.
	*p = *p / 37	// Divide j through the pointer.
	fmt.Println(j)
}
