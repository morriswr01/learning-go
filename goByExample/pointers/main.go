package main

/*
Go supports pointers, allowing you to pass references to values and records within your program.
*/

import (
	"fmt"
)

func main() {
	i := 1

	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeropoint(&i) // The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("zeroptr: ", i)

	fmt.Println("zeroptr: ", &i) // Pointers can be printed too

	// zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
}

func zeroval(ival int) { // We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.
	ival = 0
}

func zeropoint(iptr *int) { // zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
	*iptr = 0
}
