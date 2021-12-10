// Mutating Maps


// Insert or update an element in map m:

// 	example_Map[example_Key] = example_Element


// Retrieve an element:

//	example_Element = example_Map[example_Key]


// Delete and element:

//	delete(example_Map, example_Key)


// Test that a key is present with a two-value assignment:

// example_Element, ok = example_Map[example_Key]

// If key is in example_Map, ok is true.
// If not, ok is false.

// If example_ley is not in the example_map, then example_Element is the zero 
// value for the map's element type.


// If example_Element or ok have not been declared they can be using a short
// declaration form:

// 	example_Element, ok := example_Map[example_Key]

package main

import "fmt"

func main() {
	example_Map:= make(map[string]int)
	
	example_Map["Answer"] = 42
	fmt.Println("The value:", example_Map["Answer"])

	example_Map["Answer"] = 48
	fmt.Println("The value:", example_Map["Answer"])

	delete(example_Map, "Answer")
	fmt.Println("The value:", example_Map["Answer"])

	v, ok := example_Map["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
