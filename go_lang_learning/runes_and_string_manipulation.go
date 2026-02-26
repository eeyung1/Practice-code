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

//Write a function CountVowels(str string) int that counts the number of vowels (a, e, i, o, u both upper and lowercase) in a string.

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