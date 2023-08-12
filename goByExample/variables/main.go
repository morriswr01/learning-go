package main

/*
In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
*/

import "fmt"

func main() {
	var a = "initial" // var declares 1 or more variables.
	fmt.Println(a)

	var b, c int = 1,2 // You can declare multiple variables at once. with a type declaration
	fmt.Println(b,c)

	var d = true // implicit type declaration
	fmt.Println(d)

	var e int // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	// int = 0
	// bool = false
	// string = ""
	fmt.Println(e)

	f := "apple" // := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions
	fmt.Println(f)
}