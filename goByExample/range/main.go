package main

/*
range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.
*/

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for index, num := range nums { // Here we use range to sum the numbers in a slice. Arrays work like this too.
		sum += num
		fmt.Println("_", index)
	}
	fmt.Println("Sum", sum)

	// range on map iterates over key/value pairs.
	// Can also not specify v and just iterate over keys
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("$s => $s\n", k, v)
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.
	str := "string"
	for _, char := range str {
		fmt.Println(char) // Prints Unicode code
		
	}
}
