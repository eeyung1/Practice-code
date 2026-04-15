package main

import (
	"fmt"
	"os"
	"practice1/ascii"
)

func main(){
	length := len(os.Args)

	if length < 2 || length > 3 {
		fmt.Println("Usage: go run . [STRING] [banners]")
		os.Exit(1)
	}

	input := os.Args[1]

	banners := "standard"

	if length == 3 {
		banners = os.Args[2]
	}

	result, err := ascii.Render(input, banners)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}