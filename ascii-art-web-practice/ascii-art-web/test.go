package main

import (
	"ascii-art-web/ascii"
	"fmt"
)

func main(){
	result, err := ascii.Generate("Hello", "standard")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}