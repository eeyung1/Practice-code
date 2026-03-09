package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 3  {
		fmt.Println("Error reading Arguments")
		os.Exit(1)
	}

	input_text := os.Args[1]
	output_text := os.Args[2]

	data, err := os.ReadFile(input_text)
	if err != nil {
		fmt.Println("Error reading input_file")
		os.Exit(1)
	}

	text := string(data)
	result := editor(text)

	err = os.WriteFile(output_text, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
}