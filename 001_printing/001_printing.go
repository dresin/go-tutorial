package main

// there are 3 Print functions in the fmt package: Print, Println and Printf
import "fmt"

func main() {
	// Print adds blanks only around non-string items and it does not add newline
	fmt.Print("Print", 1)

	// Println always adds blanks between the items and newline at the end
	fmt.Println("Println", 1)

	// Printf provides complete control
	//Printf example without blanks and without newline
	fmt.Printf("Printf%v", 1)

	fmt.Print("Print", 2)
	fmt.Println("Println", 2)

	//Printf example without blanks and with newline
	fmt.Printf("Printf%v\n", 2)

	fmt.Print("Print", 3)
	fmt.Println("Println", 3)

	//Printf example with tab before the number and with newline
	fmt.Printf("Printf\t%v\n", 3)

	// Format verbs
	fmt.Printf("\nPrintf function uses the following verbs for formatting:\n")
	fmt.Printf("\nGeneral:\n\n")
	fmt.Printf("%%v the value in a default format\n")
	fmt.Printf("%%v+ when printing structs adds field names\n")
	fmt.Printf("%%#v a Go-syntax representation of the value\n")
	fmt.Printf("%%T a Go-syntax representation of the type\n")
	fmt.Printf("%%%% literal percent sign\n")
	fmt.Printf("\nBoolean:\n\n")
	fmt.Printf("%%t the word true or false\n")
	fmt.Printf("\nInteger:\n\n")
	fmt.Printf("%%b binary representation - base 2\n")
	fmt.Printf("%%c the character represented by the coresponding Unicode code point\n")
	fmt.Printf("%%d decimal representation - base 10\n")
	fmt.Printf("%%o octal representation - base 8\n")
	fmt.Printf("%%O octal representation with prefix 0o\n")
	fmt.Printf("%%q a single-quoted character literal safely escaped with Go syntax\n")
	fmt.Printf("%%x hexadecimal representation with lower-case letters for a-f\n")
	fmt.Printf("%%X hexadecimal representation with upper-case letters for A-F\n")
	fmt.Printf("%%U Unicode format: U+1234; same as U+%%04X\n")
	fmt.Printf("\nFloating-point and complex:\n\n")
	fmt.Printf("%%b scientific notation without decimals with the power of 2 exponent.\n")
	fmt.Printf("%%e scientific notation with lower-case e\n")
	fmt.Printf("%%E scientific notation with upper-case E\n")
	fmt.Printf("%%f decimal point without exponent\n")
	fmt.Printf("%%F synonym for %%f\n")
	fmt.Printf("%%g %%e for large exponents, %%f otherwise\n")
	fmt.Printf("%%G %%E for large exponents, %%F otherwise\n")
	fmt.Printf("%%x hexadecimal notation with power of 2 exponent\n")
	fmt.Printf("%%X upper-case hexadecimal notation\n")
	fmt.Printf("\nString and slice of bytes:\n\n")
	fmt.Printf("%%s uninterpeted bytes of string or slice\n")
	fmt.Printf("%%q a double-quoted string safely escaped with Go syntax\n")
	fmt.Printf("%%x base 16 lower-case two characters per byte\n")
	fmt.Printf("%%X base 16 upper-case two characters per byte\n")
	fmt.Printf("\nSlice:\n\n")
	fmt.Printf("%%p address of 0th element in base 16 notation with leading 0x\n")
	fmt.Printf("\nPointer:\n\n")
	fmt.Printf("%%p base 16 notation with leading 0x\n")
	fmt.Printf("The %%v, %%d, %%o, %%x and %%X verbs also work formatting the value as if it were an integer\n")

}
