package main

/*
Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

Integers and floats.
*/

import "fmt"

func main() {
	fmt.Println("GO" +  "LANG") // Strings, which can be added together with +.
	
	fmt.Println("1+1 = ", 1+1) // Integers
	fmt.Println("7.0/3.0", 7.0/3.0) // Floats
	
	fmt.Println(true && true) // Booleans
	fmt.Println(true || true)
	fmt.Println(!true)
}