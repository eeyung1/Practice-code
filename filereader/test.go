package main

import (
	"fmt"
	"strings"
)

func main() {
	// test Split — splitting by comma
	a := "he has light brain but cannot offer full nine subject"
	text := strings.Split(a,"a")
	fmt.Println("Split:", text)
	fmt.Println("Split length:", len(text))

	// test Fields — splitting by whitespace
	b := "Hello   world   this   is   a   test"
	good := strings.Fields(b)
	fmt.Println("Fields:", good)
	fmt.Println("Fields length:", len(good))

	// now see what Split does with extra spaces
	bad := strings.Split(b, " ")
	fmt.Println("Split with spaces:", bad)
	fmt.Println("Split length:", len(bad))
}


