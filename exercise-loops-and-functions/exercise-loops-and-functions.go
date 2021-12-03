package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	
	//Declare and initialize a floating point value.
	z := float64(1)
	old_Z := float64(0) 
	
		
	fmt.Printf("Calculating square root for: %v\n", x)
	
	//Repeat the calculation ten times.
	//Loop condition to stop once the value has stopped changing.
	for i := 0; i < 10 && z != old_Z; i++ {
		old_Z = z
		
		//Square root calculation.
		z -= (z*z - x) / (2*z)
		
		fmt.Printf("Round %v: Square root is: %v\n\n", i, z)
		
	}
	return z
}

func main() {
	//See how close the answer is for various values.
	for i := 0; i <10; i++ {
		fmt.Println(Sqrt(float64(i)))
	}
}
