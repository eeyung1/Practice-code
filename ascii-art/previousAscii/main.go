package main

import (
	"ascii-art/ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage: go run . \"text\" [banner]")
		os.Exit(1)
	}

	textinput := os.Args[1]
	textinput = strings.ReplaceAll(textinput, "\\n", "\n")
	banner := "standard"

	if len(os.Args) == 3 {
		banner = os.Args[2]
		if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
			fmt.Println("Use either: shadow, standard, or thinkertoy")
			os.Exit(1)
		}
	}

	result, err := ascii.Render(textinput, banner)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Print(result)
}
