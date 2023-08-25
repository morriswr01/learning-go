package main

import "fmt"

/*
Go supports methods defined on struct types.
*/

type rect struct {
	width, height int
}

func (r *rect) area() int { // This area method has a receiver type of *rect. pointer receiver
	return r.width * r.height
}

func (r rect) perim() int { // Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r // Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	r.width = 15
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
}
