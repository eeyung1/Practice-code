package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	length := len(os.Args)

	if length < 2 || length > 3 {
		fmt.Println("Error: Invalid number of Arguments")
		return
	}

	banners := "standard.txt"

	if length == 3 {
		banners = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	data, err := os.ReadFile(banners)
	if err != nil {
		fmt.Println("Error reading banner:", err)
		return
	}

	banner := strings.Split(string(data), "\n")

	lines := strings.Split(input, "\\n")

	for _, word := range lines {
		if word == ""{
			fmt.Println()
			continue
		}

		for i := 0; i <= 8; i++ {
			for j := 0; j < len(word); j++ {
				index := int(word[j]-32)*9 + i

				if index >= 0 && index < len(banner) {
					fmt.Print(banner[index])
				}
			}
			fmt.Println()
		}
	}

}
