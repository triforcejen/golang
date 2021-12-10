// Map Literals Continued

// If the top-level type is just a type name.

// It can be omitted from the elements of the literal.

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": 	{40.68433, -74.39967}, 
//	Otherwise this would read:
//	"Bell Labs": 	Vertex{40.68433, -74.39967},
	"Google":	{37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
