package main

import "fmt"

const x = 1

const (
	a = "a1"
	b = "b2"
)

func main() {
	fmt.Println("Constants are declared using the const keyword.")
	fmt.Println("Constant x is defined at the package level and initialized with the value of 1.")
	const y = 2.2
	fmt.Println("Constant y is defined at the function level and initialized with the value of 2.2.")
	fmt.Println("Constants cannot be declared using the := syntax.")
	fmt.Println("Constants can be declared together using the const() syntax.")
	fmt.Printf("Constants a and b have the values of %v and %v.\n", a, b)
	fmt.Println("Within a constant declaration, iota represents successive untyped integer constants.")
	fmt.Println("Iota starts at zero and is increased by one on every appearance.")
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Printf("c0, c1 and c2 have the values of: %v, %v and %v respectively.\n", c0, c1, c2)
	const (
		c3 = iota
		c4 = iota
		c5 = iota
	)
	fmt.Println("Iota resets to zero in every new const declaration.")
	fmt.Printf("c3, c4 and c5 have the values of: %v, %v and %v respectively.\n", c3, c4, c5)
}
