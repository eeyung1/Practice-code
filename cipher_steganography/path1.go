package main

import "fmt"

func main() {
	var b byte = 'y'
	fmt.Printf("Character: %c\n", b)
	fmt.Printf("Byte value: %d\n", b)
	fmt.Printf("Binary: %08b\n", b)
	fmt.Printf("Hexadecimal: %02x\n", b)
}