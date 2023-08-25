package main

import (
	"fmt"
	"math"
)

/*
Interfaces are named collections of method signatures.
*/

type geometry interface {
	area() float64
	perim() float64
}

type rect struct { // For our example we’ll implement this interface on rect and circle types.
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 { // To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{3, 4} // rectangle implements geometry interface
	c := circle{5}  // circle implements geometry interface

	measure(r) // The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(c)
}
