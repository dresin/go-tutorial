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
}
