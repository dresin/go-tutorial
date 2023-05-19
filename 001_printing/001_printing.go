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
	fmt.Printf("\nDefault format for %%v:\n\n")
	fmt.Printf("bool:\t\t\t\t%%t\n")
	fmt.Printf("int, int8, etc.:\t\t%%d\n")
	fmt.Printf("uint, uint8, etc.:\t\t%%d (%%#x if printed with %%#v)\n")
	fmt.Printf("float32, complex64, etc.:\t%%g\n")
	fmt.Printf("string:\t\t\t\t%%s\n")
	fmt.Printf("chan:\t\t\t\t%%p\n")
	fmt.Printf("pointer:\t\t\t%%p\n")
	fmt.Printf("\nCompund objects (recursively):\n\n")
	fmt.Printf("struct:\t\t\t{field0 field1 ...}\n")
	fmt.Printf("array, slice:\t\t[elem0 elem1 ...]\n")
	fmt.Printf("maps:\t\t\tmap[key1:value1 key2:value2 ...]\n")
	fmt.Printf("pointer to above:\t&{}, &[], &map[]\n")
	fmt.Printf("\nWidth and precision:\n\n")
	fmt.Printf("%%f\tdefault width, default precision\n")
	fmt.Printf("%%8f\twidth 8, default precision\n")
	fmt.Printf("%%.2f\tdefault width, precision 2\n")
	fmt.Printf("%%6.3f\twidth 6, precision 3\n")
	fmt.Printf("%%3.f\twidth 3, precision 0\n")
	fmt.Printf("\nOther flags:\n\n")
	fmt.Printf("+\talways print a sign (for numeric values)\n")
	fmt.Printf("\tASCII-only output for %%+q\n")
	fmt.Printf("-\tpad with spaces on the right instead left\n")
	fmt.Printf("#\talternate format:\n")
	fmt.Printf("\tadd leading 0b for binary (%%#b)\n")
	fmt.Printf("\tadd leading 0 for octal (%%#o)\n")
	fmt.Printf("\tadd leading 0x or 0X for hexadecimal (%%#x or %%#X)\n")
	fmt.Printf("\tsurpress 0x for %%#p\n")
	fmt.Printf("\t%%#q prints a raw (backquoted) string if strconv.CanBackquote returns true\n")
	fmt.Printf("\talways print a decimal point for %%#e, %%#E, %%#f, %%#F, %%#g and %%#G\n")
	fmt.Printf("\tdo not remove trailing zeros for %%#g and %%#G\n")
	fmt.Printf("\twrite  e.g. U+007a 'z' if the character is printable for %%#U\n")
	fmt.Printf("' '\tleave a space for elided sign in numbers (%% d)\n")
	fmt.Printf("\tput spaces between bytes printing strings or slices in hex (%% x, %% X)\n")
	fmt.Printf("o\tpad with leading zeros instead of spaces\n")
	fmt.Printf("\tfor numbers: move the padding after the sign\n")
	fmt.Printf("\tignored for strings, slices and byte arrays\n")
}
