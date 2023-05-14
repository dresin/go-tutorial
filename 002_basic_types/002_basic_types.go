package main

import "fmt"

func main() {
	fmt.Printf("There are six Go basic types: bool, string, int, uint, float and complex.\n\n")

	// bool
	fmt.Printf("The value %v is of type %T.\n", true, true)

	// string
	fmt.Printf("The value %v is of type %T.\n", "Go rocks!", "Go rocks!")

	// int
	fmt.Printf("The value %v is of type %T.\n", -5508, -5508)

	// uint
	fmt.Printf("The value %v is of type %T.\n", uint(2009), uint(2009))

	// float
	fmt.Printf("The value %v is of type %T.\n", -273.15, -273.15)

	// complex
	fmt.Printf("The value %v is of type %T.\n", -3.142-1.618i, -3.142-1.618i)

	fmt.Printf("\nThere is a also one pointer type: uintptr.\n\n")

	fmt.Println("The int type has four variations: int8, int16, int32 and int64.")
	fmt.Println("The uint type also has four variations: uint8, uint16, uint32 and uint64.")
	fmt.Println("There are two alias types: byte is alias for uint8, rune is alias for int32.")
	fmt.Println("The float type has two variations: float32 and float64.")
	fmt.Println("The complex type has two variations: complex64 and complex64.")

	fmt.Println()

	fmt.Println("The int, uint and uintptr types are usually 32-bit on 32-bit systems.")
	fmt.Println("The int, uint and uintptr types are usually 64-bit on 64-bit systems.")
}
