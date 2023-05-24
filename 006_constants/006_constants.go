package main

import "fmt"

const x = 1

func main() {
	fmt.Println("Constants are declared using the const keyword.")
	fmt.Println("Constant x is defined at the package level and initialized with the value of 1.")
	const y = 2.2
	fmt.Println("Constant y is defined at the function level and initialized with the value of 2.2.")
	fmt.Println("Constants cannot be declared using the := syntax.")
}
