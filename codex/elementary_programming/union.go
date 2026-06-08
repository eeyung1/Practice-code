package main 

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println()
		return
	}

	str1 := args[0]
	str2 := args[1]

	seen := make(map[byte]bool)
	result := []byte{}

	for _, ch := range str1 {
		if !seen[byte(ch)] {
			seen[byte(ch)] = true
			fmt.Println(string(ch))
			result = append(result, byte(ch))
		}
	}	
	
	for _, ch := range str2 {
		if !seen[byte(ch)] {
			fmt.Println(string(ch))
			seen[byte(ch)] = true
			result = append(result, byte(ch))
		}
	}



	fmt.Println(string(result))

}