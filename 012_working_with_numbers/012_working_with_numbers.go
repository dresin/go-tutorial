package main

import "fmt"

func main() {
	fmt.Printf("Working with numbers\n\n")
	fmt.Println("Arithmetic operators (and the types they apply to):")
	fmt.Printf("\t+\taddition\t(integer, float, complex)\n")
	fmt.Printf("\t-\tsubtraction\t(integer, float, complex)\n")
	fmt.Printf("\t*\tmultiplication\t(integer, float, complex)\n")
	fmt.Printf("\t/\tdivision\t(integer, float, complex)\n")
	fmt.Printf("\t%%\tmodulo\t\t(integer)\n\n")
	fmt.Printf("Integer overflow\n")
	fmt.Printf("\tWorking with unsigned integers using operators +, - and * may cause an overflow.\n")
	fmt.Printf("\tOverflow examples:\n")
	fmt.Printf("\t\tAdding 1 to the largest possible uint8 number using operator +:\n")
	fmt.Printf("\t\t\tvar x uint8 = 255\n")
	var x uint8 = 255
	fmt.Printf("\t\t\tfmt.Printf(\"x = %%v\", x)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("x = %v\n", x)
	fmt.Printf("\t\t\tfmt.Printf(\"x + 1 = %%v\", x + 1)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("x + 1 = %v\n", x+1)
	fmt.Printf("\t\t\t// 255 + 1 will equal to 0, and not to 256, because of the overflow\n")
	fmt.Printf("\t\t\t// the overflow happens because uint8 type has a range from 0 to 255\n")
	fmt.Printf("\t\tSubtracting 1 from the smallest possible uint8 number using operator -:\n")
	fmt.Printf("\t\t\tvar y uint8 = 0\n")
	var y uint8 = 0
	fmt.Printf("\t\t\tfmt.Printf(\"y = %%v\", y)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("y = %v\n", y)
	fmt.Printf("\t\t\tfmt.Printf(\"y - 1 = %%v\", y - 1)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("y - 1 = %v\n", y-1)
	fmt.Printf("\t\t\t// 0 - 1 will equal to 255, and not to -1, because of the overflow\n")
	fmt.Printf("\t\t\t// the overflow happens because uint8 type has a range from 0 to 255\n")
	fmt.Printf("\t\tMultiplying a uint8 number using operator *:\n")
	fmt.Printf("\t\t\tvar z uint8 = 100\n")
	var z uint8 = 100
	fmt.Printf("\t\t\tfmt.Printf(\"z = %%v\", z)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("z = %v\n", z)
	fmt.Printf("\t\t\tfmt.Printf(\"z * 3 = %%v\", z * 3)\n")
	fmt.Printf("\t\t\t// the previous line will output: ")
	fmt.Printf("z * 3 = %v\n", z*3)
	fmt.Printf("\t\t\t// 100 * 3 will equal to 44, and not to 300, because of the overflow\n")
	fmt.Printf("\t\t\t// the overflow happens because uint8 type has a range from 0 to 255\n")
}
