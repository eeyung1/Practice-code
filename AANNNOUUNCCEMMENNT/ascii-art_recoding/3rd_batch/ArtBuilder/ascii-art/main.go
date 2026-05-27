package main

import (
	"fmt"
	"builder/builder"
)

func main() {
	result := builder.NewArtBuilder().
		AddText("HI").
		SetStyle("bold").
		Build()

	fmt.Println(result)
}