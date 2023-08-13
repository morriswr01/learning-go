package main

/*
Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
*/

import (
	"fmt"
)

func main() {
	var s []string // Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // We can set and get just like with arrays.
	s[1] = "b"
	s[2] = "c"
	// s[3] = "brev" // slice is only of size 3 so we get an error here
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // len returns the length of the slice as expected.

	s = append(s, "d") // n addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	fmt.Println("sl1", l)
	fmt.Println("sl1len", len(l))

	l = append(l, "e", "f") // Slice function definitely returns a slice not an array
	fmt.Println("sl1", l)
	fmt.Println("sl1len", len(l))

	l = s[:3] // This slices up to (but excluding) s[3].
	fmt.Println("sl2", l)

	t := []string{"g", "h", "i"} // We can declare and initialize a variable for slice in a single line as well. Differece to the array is that no length is specified in the square brackets.
	fmt.Println("dcl:", t)
	fmt.Println("dcl:", append(t, "j"))

	twoD := make([][]int, 3) // Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			// twoD[i] = append(twoD[i], 69)
		}
	}
	fmt.Println("2d: ", twoD)

	// Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
}
