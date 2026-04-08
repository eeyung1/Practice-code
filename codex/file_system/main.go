package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("No file provided")
		return
	}

	input := os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content:\n", string(data))
}