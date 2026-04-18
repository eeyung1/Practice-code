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
		banners = os.Args[2]
	}

	text := os.Args[1]

	data, err := os.ReadFile(banners)
	if err != nil {
		fmt.Println("Error reading banner:", err)
		return
	}

	banner := strings.Split(string(data), "\n")
	letter := strings.Split(text, `\n`)

	for i := range banner {
		banner[i] = strings.ReplaceAll(banner[i], "\r", "")
	}

	for alphabet := range letter {
		if letter[alphabet] == "" {
			fmt.Println()
			continue
		}

		for i := 1; i < 9; i++ {
			row := []string{}

			s := letter[alphabet]

			for x := 0; x < len(s); x++ {
				index := int(s[x]-32)*9 + i
				if index >= 0 && index < len(banner) {
					row = append(row, banner[index])
				}
			}
			fmt.Println(strings.Join(row, ""))
		}
	}
}
