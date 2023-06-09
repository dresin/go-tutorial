package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Working with strings")
	fmt.Printf("\nThe length of a string can be determined using the len() function:")
	greeting := "Hello!"
	fmt.Printf("\n\tThe length of string `%s` is %d.", greeting, len(greeting))
	fmt.Printf("\n\nstrings package")
	fmt.Printf("\n\nFunction Contains determins if the string contains specific substring:")
	he := "He"
	heInGreeting := strings.Contains(greeting, he)
	fmt.Printf("\n\tDoes string `%s` contain `%s`? %t", greeting, he, heInGreeting)
	the := "The"
	theInGreeting := strings.Contains(greeting, the)
	fmt.Printf("\n\tDoes string `%s` contain `%s`? %t", greeting, the, theInGreeting)
	fmt.Printf("\n\nFunction ToUpper converts a string to uppercase:")
	fmt.Printf("\n\t%s", strings.ToUpper(greeting))
	fmt.Printf("\n\nFunction ToLower converts a string to lowercase:")
	fmt.Printf("\n\t%s", strings.ToLower(greeting))
	fmt.Printf("\n\nFunction Replace replaces a substring with another substring:")
	fmt.Printf("\n\t%s", strings.Replace(greeting, "!", " World", 1))
}
