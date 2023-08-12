package main

/*
for is Go’s only looping construct. Here are some basic types of for loops.
*/

import (
	"fmt"
)


func main() {
	i := 1
	for i <= 3 { // The most basic type, with a single condition.
		fmt.Println(i)
		i++
	}

	for j := 0; j < 7; j++ { // A classic initial/condition/after for loop.
		fmt.Println("J Loop number: ", j)
	}

	for { // for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
        fmt.Println("loop")
        break
    }

	for n := 0; n <= 5; n++ { // You can also continue to the next iteration of the loop.
		if n%2 == 0 {
            continue
        }
        fmt.Println(n)
	}
}