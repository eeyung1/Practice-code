package main

import (
	"fmt"
	"font/font"
)

func main() {
	f := font.GenerateFont()
	for _, line := range f['b'] {
		fmt.Println(line)
	}
}