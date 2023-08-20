package main

import "fmt"

/*
Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.
*/

func main() {
	a, _ := vals() // Here we use the 2 different return values from the call with multiple assignment. Must define a variable for each returned value
	fmt.Println(a)
	// fmt.Println(b)

}

func vals() (int, int) { // The (int, int) in this function signature shows that the function returns 2 ints.
	return 3, 7
}
