package main

import "fmt"

func main() {
	fmt.Println("A block is a sequence of declarations and statements within curly braces.")
	fmt.Println("There are also implicit blocks:")
	fmt.Printf("\tThe universe block encompasses all Go source text.\n")
	fmt.Printf("\tEach package has its own block.\n")
	fmt.Printf("\tEach file has a file block.\n")
	fmt.Printf("\tEach if, for and switch statement is considered to be in its own block.\n")
	fmt.Printf("\tEach clause in a switch or select statement acts like a block.\n")
	fmt.Println("Blocks can be nested.")
	fmt.Println("Blocks influence scoping.")
}
