package main

import (
	"fmt"
)

func main(){
	s := "country"

	b := []rune(s)


	fmt.Println(b)

	for i := 0; i < len(b); i+= 2 {
		b[i] = (b[i] - 32)
	}

	fmt.Println(string(b))
}

/*

r - 32 ======> lowercase ==> uppercase
r + 32 =====> uppercase ==> lower case
r - 'a' + 1 ===> letter ==> alphabet position
r - '0' ===> digit character ==> actual number

*/