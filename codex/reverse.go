package main

import (
	"fmt"
)

func main(){
	fmt.Println(reverse("awesome"))
}
func reverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	return reverse(s[1:]) + string(s[0])
}