package main

import "fmt"

const Pi = 3.14

func main(){
	const W = "World"
	fmt.Println("Hello", W)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
