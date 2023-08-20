package main

import "fmt"

/*
Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
*/

func main() {
	sum(1, 2) // Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...) // If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for num := range nums {
		total += num
	}

	fmt.Println(total)
}
