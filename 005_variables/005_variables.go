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
	fmt.Printf("\tVariable z can be initialized to %v with the following instruction:\n", z)
	fmt.Printf("\t\tz := 3\n")
	fmt.Println("Short assignment construct can be used only inside of a function.")
	fmt.Println("Once the variable is defined it's value can be changed.")
	fmt.Printf("\tThe following statement will change the value of variable z to 4:\n")
	fmt.Printf("\t\tz = 4\n")
	fmt.Println("Once the variable is defined it can not be defined again.")
	fmt.Printf("\tThe following two lines of code will cause an error when used together:\n")
	fmt.Printf("\t\tv := 1 // this line is OK\n")
	fmt.Printf("\t\tv := 2 // this line will cause an error\n")
	fmt.Println("Previously defined variable generally can not be used with := short assignment operator.")
	fmt.Printf("\tThe exception is if there is another previously not defined variable in the same statement:\n")
	fmt.Printf("\t\tw, v := 1, 2 // v can be used here only because w has never been defined before\n")
	fmt.Println("Defined variable has to be used or the code will not compile.")
	// Next line would cause an error if it weren't commented out
	// var e = 10
	fmt.Println("The only way to have a non-used variable is to name it _.")
	var _ = "This variable doesn't have to be used."
	fmt.Println("Variables declared without inital value are given zero value:")
	fmt.Printf("\t0 for numeric types,\n")
	fmt.Printf("\tfalse for the boolean type,\n")
	fmt.Printf("\t\"\" (the empty string) for strings.\n")
}
