package main

import "fmt"

func main() {
	fmt.Println("There are two forms of string literals:")
	fmt.Printf("\t1. Raw string literals\n")
	fmt.Printf("\t2. Interpreted string literals\n")
	fmt.Printf("\nRaw string literals\n")
	fmt.Printf("\tRaw string literals are defined between back quotes `\n")
	fmt.Printf("\tThey can consist of any characer except back quete.\n")
	fmt.Printf("\tThe string is composed of uninterpreted UTF-8 characters.\n")
	fmt.Printf("\tBackslashes have no special meaning.\n")
	fmt.Printf("\tString can contain newlines.\n")
	fmt.Printf("\tCarriage returns (`\\r`) are discarded.\n")
	fmt.Printf("\nInterpreted string literals\n")
	fmt.Printf("\tInterpreted string literals are defined between double quotes \"\n")
	fmt.Printf("\tThey can consist of any character except newline and unescaped double quote.\n")
	fmt.Printf("\tThree-digit octal escape (\\nnn) represents individual byte.\n")
	fmt.Printf("\tTwo-digit hexadecimal escape (\\xnn) represents individual byte.\n")
	fmt.Printf("\tAll other escapes represent possibly multi-byte UTF-8 encoded characters.\n")
}
