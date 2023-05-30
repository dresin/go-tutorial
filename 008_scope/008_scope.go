package main

import "fmt"

func main() {
	fmt.Println("A declaration binds a non-blank identifier to a specific scope.")
	fmt.Println("No identifier may be declared more than once in the same block.")
	fmt.Println("No identifier may be declared in both the package and file block.")
	fmt.Println("The blank identifier does not introduce a binding.")
	fmt.Println("The init identifier may only be used for init function declarations in the package block.")
	fmt.Println("The init identifier does not introduce a new binding.")
}
