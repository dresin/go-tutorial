package main

import "fmt"

func main() {
	fmt.Println("There are six Go basic types: bool, string, int, uint, float and complex")

	// bool
	fmt.Printf("The value %v is of type %T\n", true, true)

	// string
	fmt.Printf("The value %v is of type %T\n", "Go rocks!", "Go rocks!")

	// int
	fmt.Printf("The value %v is of type %T\n", -5508, -5508)

	// uint
	fmt.Printf("The value %v is of type %T\n", uint(2009), uint(2009))

	// float
	fmt.Printf("The value %v is of type %T\n", -273.15, -273.15)

	// complex
	fmt.Printf("The value %v is of type %T\n", -3.142-1.618i, -3.142-1.618i)
}
