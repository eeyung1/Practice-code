package main

import (
	"fmt"
	"generator/generator"
)

func main() {
	for _, line := range generator.GeneratePattern('Z') {
		fmt.Println(line)
	}
}