package main

import "fmt"

func main() {
	fmt.Printf("Escapes:\n\n")
	fmt.Printf("\\a\tU+0007 alert or bell\n")
	fmt.Printf("\\b\tU+0008 backspace\n")
	fmt.Printf("\\f\tU+000C form feed\n")
	fmt.Printf("\\n\tU+000A line feed or new line\n")
	fmt.Printf("\\r\tU+000D carriage return\n")
	fmt.Printf("\\t\tU+0009 horizontal tab\n")
	fmt.Printf("\\v\tU+000B vertical tab\n")
	fmt.Printf("\\\\\tU+005C backslash\n")
	fmt.Printf("\\'\tU+0027 single quote (only with rune literals)\n")
	fmt.Printf("\\\"\tU+0022 double quote (only with string literals)\n")
}
