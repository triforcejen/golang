package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	//Declare and initialize a floating point value
	z := float64(1)
	fmt.Printf("Calculating square root for: %v\n", x)
	//Repeat the calculation ten times
	for i := 0; i < 10; i++ {
		
		//Square root calculation
		z -= (z*z - x) / (2*z)
		fmt.Printf("Round %v: Square root is: %v\n", i, z)
	}
	return z
}

func main() {
	for i := 0; i <10; i++ {
		fmt.Println(Sqrt(float64(i)))
	}
}
