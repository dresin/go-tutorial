package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Working with strings")
	fmt.Printf("\nThe length of a string can be determined using the len() function:")
	fmt.Printf("\n\tThe length of string `Hello!` is %d.", len("Hello!"))
	fmt.Printf("\n\tlen(\"Hello!\") will return:")
	fmt.Printf("\n\t\t%d", len("Hello!"))
	fmt.Printf("\n\nstrings package")
	fmt.Printf("\n\nFunction Contains determines if the string contains specific substring:")
	fmt.Printf("\n\tDoes string `Hello!` contain `He`? %t", strings.Contains("Hello!", "He"))
	fmt.Printf("\n\tstrings.Contains(\"Hello!\", \"He\") will return:")
	fmt.Printf("\n\t\t%t", strings.Contains("Hello!", "He"))
	fmt.Printf("\n\tDoes string `Hello!` contain `The`? %t", strings.Contains("Hello!", "The"))
	fmt.Printf("\n\tstrings.Contains(\"Hello!\", \"The\") will return:")
	fmt.Printf("\n\t\t%t", strings.Contains("Hello!", "The"))
	fmt.Printf("\n\nFunction ToUpper converts a string to uppercase:")
	fmt.Printf("\n\tstrings.ToUpper(\"Hello!\") will return:")
	fmt.Printf("\n\t\t%s", strings.ToUpper("Hello!"))
	fmt.Printf("\n\nFunction ToLower converts a string to lowercase:")
	fmt.Printf("\n\tstrings.ToLower(\"Hello!\") will return:")
	fmt.Printf("\n\t\t%s", strings.ToLower("Hello!"))
	fmt.Printf("\n\nFunction Replace replaces a substring with another substring:")
	fmt.Printf("\n\tstrings.Replace(\"Hello!\", \"!\", \" World\") will return:")
	fmt.Printf("\n\t\t%s", strings.Replace("Hello!", "!", " World", 1))
	fmt.Printf("\n\nFunction Compare lexicographically compares two strings:")
	fmt.Printf("\n\tIt returns an integer -1, 0 or 1")
	fmt.Printf("\n\tstrings.Compare(string1, string2) will return:")
	fmt.Printf("\n\t\t-1 if string1 contains unicode characters that come before string2 characters,")
	fmt.Printf("\n\t\t 0 if string1 and string2 are equal,")
	fmt.Printf("\n\t\t 1 if string1 contains unicode characters that come after string2 characters.")
	fmt.Printf("\n\tThe unicode order of characters is evaluated from left to right, resprectively.")
	s1 := "aaa"
	s2 := "bb"
	s3 := "bb"
	s4 := "a"
	fmt.Printf("\n\t%d", strings.Compare(s1, s2))
	fmt.Printf("\n\t%d", strings.Compare(s2, s3))
	fmt.Printf("\n\t%d", strings.Compare(s3, s4))
	fmt.Printf("\n\nFunction Split splits the string into substrings:")
	fmt.Printf("\n\tstrings.Split(\"one,two\", \",\") will return:")
	fmt.Printf("\n\t\t%s", strings.Split("one,two", ","))
	fmt.Printf("\n\tstring \"one,two\" is split into substrings \"one\" and \"two\" using \",\" as a separator.")
	fmt.Printf("\n\nFunction Count counts the number of non-overlaping instances of substring")
	fmt.Printf("\n\tstrings.Count(\"Hello\", \"l\") will return:")
	fmt.Printf("\n\t\t%d", strings.Count("Hello", "l"))
	fmt.Printf("\n\nFunction Repeat returns a string repeated given number of times:")
	fmt.Printf("\n\tstrings.Repeat(\"Hey\", 3) will return:")
	fmt.Printf("\n\t\t%s", strings.Repeat("Hey", 3))
	fmt.Printf("\n\nFunction Index returns the index of the first substring appearance:")
	fmt.Printf("\n\tCharacters are enumerated using indices starting from 0.")
	fmt.Printf("\n\tThe first character in any string has index 0, the second character has index 1 and so on.")
	fmt.Printf("\n\tThe last character in any string has an index that is 1 less than the number of characters that the string has.")
	fmt.Printf("\n\tstrings.Index.(\"Hello\", \"l\") will return:")
	fmt.Printf("\n\t\t%d", strings.Index("Hello", "l"))
}
