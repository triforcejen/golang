// Exercise: Errors

// Copy your Sqrt function from the earlier exercise and modify it to return an error value.


// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.


// Create a new type:

	// type ErrNegativeSqrt float64

// and make it an error by giving it a:

	// func (e ErrNegativeSqrt) Error() string

// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".


// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.

// You can avoid this by converting e first:

	// fmt.Sprint(float64(e)).

// Why?


// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

//Declare and initialize a floating point value.


// This link really helps: 

// https://rmoff.net/2020/07/01/learning-golang-some-rough-notes-s01e06-errors/


package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	
	if(x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	
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
	return z, nil
}

func main() {
	for _, x := range []float64{-9, 9} {
		if result, error_Found := Sqrt(x); error_Found == nil {
			fmt.Printf("Result: %v\n", result)
		} else {
			fmt.Printf("ERROR: %v\n", error_Found.Error())
		}
	}
}
