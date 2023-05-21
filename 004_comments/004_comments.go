package main

import "fmt"

func main() {
	fmt.Println("There are two types of comments in Go:")
	fmt.Printf("\t1. Line comments which start with // and stop at the end of the line,\n")
	fmt.Printf("\t2. General comments which start with /* and end with (the first) */.\n")
	fmt.Printf("\nA comment can not start inside of a rune or string litteral.\n")
	fmt.Printf("A comment can not start inside of a comment.\n")
	fmt.Printf("A general comment containing no new lines acts like space.\n")
	fmt.Printf("Any other comment acts like a new line.")
}
