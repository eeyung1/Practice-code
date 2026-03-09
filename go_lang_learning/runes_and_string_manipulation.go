// package main

// import "fmt"

// func main() {
//     name := "Hello"
//     fmt.Println(len(name))
// }


// package main

// import "fmt"

// func main() {
//     name := "Héllo"
//     fmt.Println(len(name))
// }


// What is a Rune?
// A rune is Go's way of representing a single character — properly, regardless of how many bytes it takes up.

// package main

// import "fmt"

// func main() {
//     name := "Héllo"
//     fmt.Println(len(name))                    // bytes
//     fmt.Println(len([]rune(name)))            // actual characters
// }


// package main

// import "fmt"

// func main() {
//     name := "Hello"
//     for i, ch := range name {
//         fmt.Println(i, ch)
//     }
// }

// package main

// import "fmt"

// func main() {
//     name := "Hello"
//     for i, ch := range name {
//         fmt.Println(i, string(ch))
//     }
// }


// package main

// import "fmt"

// func main() {
//     word := "Hello1"
//     for _, ch := range word {

// 		if ch >= '0' && ch <= '9' {
// 			fmt.Println("Number found!")
// 		}

//     }
// }

// package main

// import "fmt"

// func CheckNumber(arg string) bool {
// 	for _, r := range arg {
// 		if r >= '0' && r <= '9' {
// 			return true
// 		}

// 	}
// 			return false
// }


// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }


// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func CheckNumber(arg string) bool {
// 	for _, r := range arg {
// 		if unicode.IsDigit(r) {
// 			return true
// 		}
// 	}
// 			return false
// }


// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }



// package main

// import "fmt"

// func CountAlpha(arg string) int{
// 	count := 0

// 	for _, r := range arg {
// 		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }


// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func CountAlpha(arg string) int{
// 	count := 0

// 	for _, r := range arg {
// 		if unicode.IsLetter(r) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }

// The unicode package has a whole family of these handy functions:

// unicode.IsLetter(r) → is it a letter?
// unicode.IsDigit(r) → is it a digit? (replaces your r >= '0' && r <= '9')
// unicode.IsUpper(r) → is it uppercase?
// unicode.IsLower(r) → is it lowercase?
// unicode.IsSpace(r) → is it a space, tab, newline?

// package main

// import ("fmt")

// func CountChar(str string, c rune) int {
// 	count := 0

// 	for _, r := range str {
// 		if r == c {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountChar("Hello World", 'l'))
// 	fmt.Println(CountChar("5  balloons", '5'))
// 	fmt.Println(CountChar("   ", ' '))
// 	fmt.Println(CountChar("The 7 deadly sins", '7'))
// }

// package main

// import ("fmt")

// func PrintIf(str string) string {
// 	if str == "" || len(str) >= 3 {
// 		return "G\n"
// 	}
// 	return "Invalid Input\n"
// }

// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }

// package main

// import ("fmt")

// func PrintIfNot(str string) string {
// 	if len(str) < 3 {
// 		return "G\n"
// 	}
// 	return "Invalid Input\n"
// }


// func main() {
// 	fmt.Print(PrintIfNot("abcdefz"))
// 	fmt.Print(PrintIfNot("abc"))
// 	fmt.Print(PrintIfNot(""))
// 	fmt.Print(PrintIfNot("14"))
// }


// package main

// import ("fmt")

// func RectPerimeter(w, h int) int {
// 	if w < 0 || h < 0 {
// 		return -1
// 	}
// 	return 2*(w+h)
// }


// func main() {
// 	fmt.Println(RectPerimeter(10, 2))
// 	fmt.Println(RectPerimeter(434343, 898989))
// 	fmt.Println(RectPerimeter(10, -2))
// }

// exercise 1.

// Write a function CountVowels(str string) int that counts the number of vowels (a, e, i, o, u both upper and lowercase) in a string.

// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func CountVowels(str string) int {
// 	count := 0

// 	for _, r := range str {
// 		if unicode.IsLetter(r) && (r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u') || unicode.IsLetter(r) && (r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U') {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main(){
// 	fmt.Println(CountVowels("Aeghidho"))
// 	fmt.Println(CountVowels("Hello World"))

// }

// optimized version

// func CountVowels(str string) int {
//     count := 0
//     for _, r := range str {
//         r = unicode.ToLower(r)
//         if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
//             count++
//         }
//     }
//     return count
// }

// package main

// import ("fmt")

// func main(){
// 	str := "Hello World"
// 	fmt.Println(str[0:5])
// }

// package main

// import "fmt"

// func RetainFirstHalf(str string) string {
//     if str == "" {
//         return ""
//     } else if len(str) == 1 {
//         return str
//     }

//     half := len(str) / 2
//     return str[half:]
// }

// func main() {
// 	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
// 	fmt.Println(RetainFirstHalf("A"))
// 	fmt.Println(RetainFirstHalf(""))
// 	fmt.Println(RetainFirstHalf("Hello World"))
// }

// exercise 2. 

// Write a function IsPalindrome(str string) bool that returns true if a string reads the same forwards and backwards.

// package main

// import "fmt"

// func IsPalindrome(str string) bool {
// 	left := 0
// 	right := len(str)-1

// 	for left <= right {
// 		if str[left] != str[right] {
// 			return false
// 		}

// 		left ++
// 		right --
// 	}

// 	return true
// }

// func main(){
// 	fmt.Println(IsPalindrome("racecar"))
// 	fmt.Println(IsPalindrome(""))
// 	fmt.Println(IsPalindrome("hello"))
// 	fmt.Println(IsPalindrome("A"))
// }

// package main

// import (
// 	"fmt")

// func CamelToSnakeCase(str string) string {
// 	if str == "" {
// 		return ""
// 	}

// 	result := ""
// 	prevUpper := false

// 	for i, r := range str {
// 		isLower, isUpper := (r >= 'a' && r <= 'z'), (r >= 'A' && r <= 'Z')

// 		if !isLower && !isUpper || isUpper && prevUpper {
// 			return str
// 		}

// 		if isUpper && i != 0 {
// 			result += "_"
// 		}

// 		result += string(r)
// 		prevUpper = isUpper
// 	}

// 	return result

// }


// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }

// package main

// import (
// 	"fmt"
// )

// func DigitLen(n, base int) int {
// 	if base < 2 || base > 36 {
// 		return -1
// 	}

// 	if n < 0 {
// 		n = -n
// 	}

// 	count := 0

// 	for n > 0 {
// 		n = n/base
// 		count++
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(DigitLen(100, 10))
// 	fmt.Println(DigitLen(100, 2))
// 	fmt.Println(DigitLen(-100, 16))
// 	fmt.Println(DigitLen(100, -1))
// }

// package main

// import (
//     "fmt"
// 	"strings"
// )

// func FirstWord(s string) string {
// 	if s == "" {
// 		return ""
// 	}

// 	words := strings.Fields(s)
// 	if len(words) == 0 {
// 		return "\n"
// 	}

// 	//first word
// 	//return words[0] + "\n"

// 	//lastword
// 	return words[len(words)-1] + "\n"

// }

// func main() {
//     fmt.Print(FirstWord("hello there"))
//     fmt.Print(FirstWord("how are you doing over there now"))
//     fmt.Print(FirstWord("hello   .........  bye"))
// }


// package main

// import (
// 	"fmt"
// )

// func FishAndChips(n int) string {
// 	if n < 0 {
// 		return "error: number is negative"
// 	}

// 	if n%2 == 0 && n%3 == 0 {
// 		return "fish and chips"
// 	} else if n%2 == 0 {
// 		return "fish"
// 	} else if n%3 == 0 {
// 		return "chips"
// 	} else {
// 		return "error: non divisible"
// 	}
// }

// func main() {
// 	fmt.Println(FishAndChips(4))
// 	fmt.Println(FishAndChips(9))
// 	fmt.Println(FishAndChips(6))
// }

// package main

// import (
// 	"fmt"
// )

// func Gcd(a, b uint) uint {
// 	if a == 0 || b == 0 {
// 		return 0
// 	}

// 	for b != 0 {
// 		remainder := a%b
// 		a = b
// 		b = remainder
// 	}

// 	return a
// }

// func main() {
// 	fmt.Println(Gcd(42, 10))
// 	fmt.Println(Gcd(42, 12))
// 	fmt.Println(Gcd(14, 77))
// 	fmt.Println(Gcd(17, 3))
// }

//function to reverse string

// package main

// import (
// 	"fmt"
// )

// func reverseString(arg string) string {
// 	result := ""
// 	for _, r := range arg {
// 		result = string(r) + result 
// 	}

// 	return result 
// }

// func main(){
// 	fmt.Println(reverseString("Hello"))
// }

package main

import (
	"fmt"
)

func HashCode(dec string) string {
	length := len(dec)
	result := ""

	for _, r := range dec {
		ascii := int(r)
		hashed := (ascii + length) % 127
		
		if hashed < 32 {
			hashed += 33
		}

		result += string(rune(hashed))
	}

	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}