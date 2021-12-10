// Exercise: Slices

////////////////////////////////////////////////////////////////////////////////
// WARNING: YOU NEED TO GO TO THE FOLLOWING SITE TO RUN THIS CODE:            //
// https://go.dev/tour/moretypes/18				              //
////////////////////////////////////////////////////////////////////////////////

// Implement Pic.

// It should return a slice of length dy.

// Each element of which is a slice of dx 8-bit unsigned integers.

// When you run the program it will display your picture, interpreting the
// integers as grayscale values.

// The choice of image is up to you.

// Interesting functions include (x + y), x * y, and x ^ y.

// You need to use a loop to allocate each []uint8 inside the [][]unit8.

// Use uint8(intValue) to convert between types.

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	new_Pic := make([][]uint8, dy)
	
	for i := range new_Pic{
		new_Pic[i] = make([]uint8, dx)
	}
	for i := range new_Pic{
			for j := range new_Pic[i]{
				new_Pic[i][j] = uint8(i * j)
			} 
		}
	return new_Pic
}

func main() {
	pic.Show(Pic)
}
