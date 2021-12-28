// Function Values:

// Functions are values too.

// They can be passed around just like other values.

// Function values may be used as function arguments and return values.

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypotenuse(5, 12))

	fmt.Println(compute(hypotenuse))
	fmt.Println(compute(math.Pow))
}
