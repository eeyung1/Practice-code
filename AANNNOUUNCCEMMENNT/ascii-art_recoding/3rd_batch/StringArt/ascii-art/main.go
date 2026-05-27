package main

import (
	"fmt"
	"converter/converter"
)

func main() {
	fmt.Println(converter.StringToArt("1"))
	fmt.Println(converter.StringToArt("12"))
	fmt.Println(converter.StringToArt("1\n2"))
}