package main

import (
	"os"
	"fmt"
)

func main(){
	args := len(os.Args)

	if args < 2 || args > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	input := os.Args[1]
	banner := "standard"

	if args == 3 {
		banner = os.Args[2]
	}

	bannerMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := Render(input, bannerMap)

	fmt.Print(result)
}