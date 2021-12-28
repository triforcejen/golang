// Exercise: Maps

// Implement WordCount.

// It should return a map of the counts of each "word" in the string 
// s.

// The wc.Test function runs a test suite against the provided function and
// prints success or failure.

// Steps:
// 1: Create slice containing each word in the string using strings.Fields(s)
// 2: Create a map with each word in the slice as the key with the value being 0.
// 3: For each key word in the map:
// 4:	Iterate through the slice and count how many times the key word is in the slice.
// 5:	Add the count as the value for that key word in the map.

package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func Split_String(sentence string) []string {
		
	// Create empty slice.
	var words []string

	// Append words to the slice.
	for _, word := range strings.Fields(sentence){
		words = append(words, word)
	}

	fmt.Println(words)
	return words
}

func Map_Words(words []string) map[string]int {
	
	// Create empty map.
	word_Map := make(map[string]int)

	// Append words in slice to keys.
	// Set values to 0.
	for _, word := range words{
		word_Map[word] = 0
	}

	return word_Map
}

func Map_Occurences(words []string, word_Map map[string]int) map[string]int {
	// For each key in map.
	for key, _ := range word_Map {
		// For each word in slice.
		for _, word := range words {
			if word == key {
				fmt.Println("Key: ", key, " matches Word: ", word)
				word_Map[key]++
			} else {
				fmt.Println("Key: ", key, " !matches Word: ", word)
			}
		}
	 }

	return word_Map
}

func Word_Count_Test(){
	var test_Sentence string = "I am learning Go!"
	var test_Slice []string
	test_Map := make(map[string]int)
	
	fmt.Printf("\n")

	fmt.Println("Testing string splitting into slice of seperate words:")
	
	fmt.Println("Input is:")
	fmt.Println(test_Sentence)
	
	fmt.Println("Calling function...")
	test_Slice = Split_String(test_Sentence)
	
	fmt.Println("Output is:")
	fmt.Println(test_Slice)
	for _, word := range test_Slice{
		fmt.Println(word)
	}
	
	fmt.Printf("\n")

	fmt.Println("Testing mapping slice of seperate words:")
	
	fmt.Println("Input is:")
	fmt.Println(test_Slice)

	fmt.Println("Calling function...")
	test_Map = Map_Words(test_Slice)

	fmt.Println("Output is:")
	fmt.Println(test_Map)

	fmt.Printf("\n")

	fmt.Println("Testing mapping occurences of key in slice as values for the keys")
	
	fmt.Println("Input is:")
	fmt.Println(test_Slice)
	fmt.Println(test_Map)

	fmt.Println("Calling function...")
	test_Map = Map_Occurences(test_Slice, test_Map)

	fmt.Println("Output is:")
	fmt.Println(test_Map)

	fmt.Printf("\n")

	fmt.Println("End of tests")
	fmt.Printf("\n")
}

func WordCount(s string) map[string]int {
	return Map_Occurences(Split_String(s), Map_Words(Split_String(s)))
}

func main() {
//	Word_Count_Test()
	wc.Test(WordCount)
}
