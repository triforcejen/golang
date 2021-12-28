// Exercise: Fibonacci Closure

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers 
// (0, 1, 1, 2, 3, 5, ...)

package main

import "fmt"

// Fibonacci is a function that returns the sum of the two numbers that precede it.

// 1: 0  + 1  = 1
// 2: 1  + 1  = 2
// 3: 1  + 2  = 3
// 4: 2  + 3  = 5
// 5: 3  + 5  = 8
// 6: 5  + 8  = 13
// 7: 8  + 13 = 21
// 8: 13 + 21 = 34
// 9: 21 + 34 = 55

func fibonacci() func() int {
	first_Value := 0
	second_Value := 1
	return func() int {
		sum := first_Value + second_Value
		first_Value = second_Value
		second_Value = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
