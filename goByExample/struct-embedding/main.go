package main

import "fmt"

/*
Go supports embedding of structs and interfaces to express a more seamless composition of types. This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.
*/

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %v", b.num)
}

type container struct {
	base // A container embeds a base. An embedding looks like a field without a name.
	str  string
}

func main() {
	co := container{
		base: base{ // When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // We can access the base’s fields directly on co, e.g. co.num.
	fmt.Println("also num:", co.base.num)                 // Alternatively, we can spell out the full path using the embedded type name.

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("desciber:", d.describe()) // Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base.
}
