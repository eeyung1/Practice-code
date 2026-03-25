package main

import (
	"fmt"
)

// 1. Takes a string, returns its length as an int
func stringLength(s string) int {
	return len(s)
}

// 2. Takes two ints, returns their sum AND their product
func sumAndProduct(a int, b int) (int, int) {
	return (a + b), (a * b)
}

// 3. Takes a string and an int index, returns the character at
//    that index as a rune AND an error if index is out of bounds
func charAtIndex(s string, index int) (rune, error) {
	return rune(s[index]), nil
}

// 4. Takes a string, returns true if it contains only
//    printable ASCII characters (codes 32–126), false otherwise
func isAllPrintableASCII(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(stringLength("awesome"))
	fmt.Println(sumAndProduct(3, 4))
	fmt.Println(charAtIndex("awesome", 4))
	fmt.Println(isAllPrintableASCII("awesome"))
}