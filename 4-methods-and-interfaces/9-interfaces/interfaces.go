// Interfaces:

// An interface type is defined as a set of method signatures.

// A value of interface type can hold any value that implements those methods.

// Note: There is an error in the example code on line 22.
// Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex 
// (the pointer type).

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main(){
	var a Abser
	f := Example_Float(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f	// A Example_Float implements Abser.
	a = &v	// A *Vertex implements Abser.

	// In the following line, v is a Vertex.
	// (Not a *Vertex).
	// And does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type Example_Float float64

func (f Example_Float) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}