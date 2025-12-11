package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	source := os.Args[1]
	oldcharacter := os.Args[2]
	newcharacter := os.Args[3]

	if len(oldcharacter) != 1 || len(newcharacter) != 1 {
		return
	}

	result := ""
	for _, y := range source {
		if string(y) == oldcharacter {
			result += newcharacter
		} else {
			result += string(y)
		}
	}

	fmt.Println(result)
}
