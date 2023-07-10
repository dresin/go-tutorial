package main

import "fmt"

func main() {
	fmt.Println("Working with strings")
	fmt.Printf("\nThe length of a string can be determined using the len() function:")
	greeting := "Hello!"
	fmt.Printf("\n\tThe length of string `%s` is %d.", greeting, len(greeting))
}
