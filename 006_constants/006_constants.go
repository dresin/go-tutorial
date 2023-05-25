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
	fmt.Printf("Constants a and b have the values of %v and %v.", a, b)
}
