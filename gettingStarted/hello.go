package main

import "fmt"
import "rsc.io/quote" // Auto installed this using `go mod tidy`

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())
}