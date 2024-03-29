package main

import "fmt"

func main() {
	fmt.Println("There are three numeric types in Go:")
	fmt.Printf("\tinteger\n")
	fmt.Printf("\tfloating-point\n")
	fmt.Printf("\tcomplex\n\n")
	fmt.Println("The integer type has two variations:")
	fmt.Printf("\tunsigned integer\n")
	fmt.Printf("\tsigned integer\n\n")
	fmt.Println("Unsigned integers have four variations:")
	fmt.Printf("\tuint8\t(0 to 255)\n")
	fmt.Printf("\tuint16\t(0 to 65535)\n")
	fmt.Printf("\tuint32\t(0 to 4294967295)\n")
	fmt.Printf("\tuint64\t(0 to 18446744073709551615)\n\n")
	fmt.Println("Signed integers have four variations:")
	fmt.Printf("\tint8\t(-128 to 127)\n")
	fmt.Printf("\tint16\t(-32768 to 32767)\n")
	fmt.Printf("\tint32\t(-2147483648 to 2147483647)\n")
	fmt.Printf("\tint64\t(-9223372036854775808 to 9223372036854775807)\n\n")
	fmt.Println("Float has two variations:")
	fmt.Printf("\tfloat32\t(IEEE-754 32-bit floating-point numbers)\n")
	fmt.Printf("\tfloat64\t(IEEE-754 64-bit floating-point numbers)\n\n")
	fmt.Println("There are two numeric types used as aliases:")
	fmt.Printf("\tbyte\t(alias for uint8)\n")
	fmt.Printf("\trune\t(alias for int32, represents a Unicode code point)\n\n")
	fmt.Println("Integer types with implementation-specific sizes:")
	fmt.Printf("\tuint\t(32 bits on 32-bit systems / 64 bits on 64-bit systems)\n")
	fmt.Printf("\tint\t(32 bits on 32-bit systems / 64 bits on 64-bit systems)\n")
	fmt.Printf("\tuintptr\t(uint large enough to store a pointer value)\n\n")
}
