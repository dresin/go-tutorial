package main

import "fmt"

var x = 1

func main() {
	fmt.Println("The variables are declared with the var statement.")
	fmt.Printf("Variable x is declared at the package level with the initial value of %v.\n", x)
	var y = 2
	fmt.Printf("Variable y is declared at the function level with the initial value of %v.\n", y)
	z := 3
	fmt.Println("Variable z is declared using the short assignment := instead of using var.")
	fmt.Printf("Variable z has the initial value of %v.\n", z)
	fmt.Println("Short assignment construct can be used only inside of a function.")
	fmt.Println("Defined variable has to be used or the code will not compile.")
	// Next line would cause an error if it weren't commented out
	// var e = 10
	fmt.Println("The only way to have a non-used variable is to name it _.")
	var _ = "This variable doesn't have to be used."
}
