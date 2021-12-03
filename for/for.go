package main

import "fmt"

func main(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i: %v \t sum: %v\n",i, sum)
	}
	fmt.Println(sum)
}
