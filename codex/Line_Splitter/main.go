package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("No file provided")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	clean := strings.TrimSpace(string(data))
	lines := strings.Split(clean, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		fmt.Printf("%d: %s\n", i, line)
	}
}