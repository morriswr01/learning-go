package main

/*
In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.
*/

import (
	"fmt"
)

func main() {
	var a [5]int // Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
	fmt.Println("empty:", a)

	a[4] = 100 // We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a)) // The builtin len returns the length of an array.

	b := [5]int{1, 2, 3, 4, 5} // Use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	var twoD [2][3]int // Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
