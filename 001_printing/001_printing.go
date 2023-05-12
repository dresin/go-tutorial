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
}
