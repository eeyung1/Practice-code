package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, r := range arg {
		if r >= '0' && r <= '9' {
	return true
	}
  }
return false
   }


func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
