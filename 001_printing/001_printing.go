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
	fmt.Println("%v the value in a default format")
	fmt.Println("%v+ when printing structs adds field names")
	fmt.Println("%#v a Go-syntax representation of the value")
	fmt.Println("%T a Go-syntax representation of the type")
	fmt.Println("%% literal percent sign")
	fmt.Printf("\nBoolean:\n\n")
	fmt.Println("%t the word true or false")
	fmt.Printf("\nInteger:\n\n")
	fmt.Println("%b binary representation - base 2")
	fmt.Println("%c the character represented by the coresponding Unicode code point")
	fmt.Println("%d decimal representation - base 10")
	fmt.Println("%o octal representation - base 8")
	fmt.Println("%O octal representation with prefix 0o")
	fmt.Println("%q a single-quoted character literal safely escaped with Go syntax")
	fmt.Println("%x hexadecimal representation with lower-case letters for a-f")
	fmt.Println("%X hexadecimal representation with upper-case letters for A-F")
	fmt.Println("%U Unicode format: U+1234; same as U+%04X")
}
