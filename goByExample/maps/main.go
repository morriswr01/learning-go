package main

/*
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages)..
*/

import (
	"fmt"
)

func main() {
	m := make(map[string]int) // To create an empty map, use the builtin make: make(map[key-type]val-type).

	m["k1"] = 7 // Set key/value pairs using typical name[key] = val syntax
	m["k2"] = 13

	fmt.Println("map:", m) // Printing a map with e.g. fmt.Println will show all of its key/value pairs. Eg map: map[k1:7 k2:13]

	v1 := m["k1"] // Get a value for a key with name[key].
	v3 := m["k3"] // If the key doesn’t exist, the zero value of the value type is returned.
	fmt.Println("v1", v1)
	fmt.Println("v3", v3)

	fmt.Println("len:", len(m)) // The builtin len returns the number of key/value pairs when called on a map.

	delete(m, "k2") // The builtin delete removes key/value pairs from a map.
	fmt.Println("map:", m["k2"])

	key, prs := m["k2"] // The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _. I have switched this _ to "key"
	fmt.Println("key", key)
	fmt.Println("present", prs)

	n := map[string]int{"foo": 1, "bar": 2} // You can also declare and initialize a new map in the same line with this syntax. No need for make here so theoretically you can do map[string]int{} to create an empty map
	fmt.Println("map:", n)

	
}
