package main

import (
	//"fmt"
	"unicode"
)

func toUpper(s string) string {
	
	b := []rune(s)

	for i := 0; i < len(b); i+= 2 {
		b[i] = (b[i] - 32)
	}

	return string(b)
}

func changeStr(b string) string {
	word := ""
	for i, r := range b {
		if i % 2 == 0 && unicode.IsLetter(r) {
			r -= 32
		}

		word += string(r)
	}

	return word
}

func firstChar(b string) string {
	word := ""
	for i, r := range b {
		if i == 0 && unicode.IsLetter(r) {
			r -= 32
		}

		word += string(r)
	}

	return word
}

func lastChar(b string) string {
	word := ""
	for i, r := range b {
		if i == len(b) - 1 && unicode.IsLetter(r) {
			r -= 32
		}

		word += string(r)
	}

	return word
}


//rune manipulation

package main

import (
	"fmt"
)

func main() {
	for r := 65; r <= 90; r++ {
		if r % 2 == 0 {
			fmt.Printf("%c ", r + 32)
		} else {
			fmt.Printf("%c ", r)
		}
	}
}


// func thirdTimeIsCharm(b string) string {
// 	word := ""
// 	// for i, r := range b {

// 	// 	return b[1:2] + "go"
// 	// 	if i % 2 == 0 {
// 	// 		word += string(r)
// 	// 	}
// 	// }

// 	return word
// }

// func main(){
// 	fmt.Println(toUpper("country"))
// 	fmt.Println(changeStr("country"))
// 	fmt.Println(firstChar("hello"))
// 	fmt.Println(lastChar("awesome"))
// 	fmt.Println(thirdTimeIsCharm("123456789"))
// }

/*

r - 32 ======> lowercase ==> uppercase
r + 32 =====> uppercase ==> lower case
r - 'a' + 1 ===> letter ==> alphabet position
r - '0' ===> digit character ==> actual number

*/