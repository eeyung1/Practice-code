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

// package main

// import (
// 	"fmt"
// )

// func HashCode(dec string) string {
// 	length := len(dec)
// 	result := ""

// 	for _, r := range dec {
// 		ascii := int(r)
// 		hashed := (ascii + length) % 127

// 		if hashed < 32 {
// 			hashed += 33
// 		}

// 		result += string(rune(hashed))
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func RepeatAlpha(s string) string {
// 	result := ""

// 	for _, r := range s {
// 		if unicode.IsLower(r) {
// 			index := r - 'a' + 1
// 			result += strings.Repeat(string(r), int(index))
// 		} else if unicode.IsUpper(r) {
// 			index := r - 'A' + 1
// 			result += strings.Repeat(string(r), int(index))
// 		} else {
// 			result += string(r)
// 		}
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println(RepeatAlpha("abc"))
// 	fmt.Println(RepeatAlpha("Choumi."))
// 	fmt.Println(RepeatAlpha(""))
// 	fmt.Println(RepeatAlpha("abacadaba 01!"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"os"
// )

// func main(){
// 	if len(os.Args) != 4 {
// 		fmt.Println("Error reading CLI")
// 		return
// 	}

// 	text := os.Args[1]
// 	old := os.Args[2]
// 	new := os.Args[3]

// 	result := strings.ReplaceAll(text, old, new)

// 	fmt.Println(result)

// 	// result := ""

// 	// for _, r := range text {
// 	// 	if r == rune(old[0]) {
// 	// 		result += string(new[0])
// 	// 	} else {
// 	// 		result += string(r)
// 	// 	}
// 	// }

// 	// fmt.Print(result + "\n")
// }

// //TEXTFORMATTER

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"strconv"
// 	"unicode"
// )

// func TextFormatter(text string) string {
// 	word := strings.Fields(text)

// 	for i := 0; i < len(word); i++ {
// 		if word[i] == "(hex)" {
// 			data, err := strconv.ParseInt(word[i-1], 16, 64)
// 			if err != nil {
// 				return "Error"
// 			}

// 			word[i-1] = strconv.FormatInt(data, 10)

// 			word = append(word[:i], word[i+1:]...)
// 			i--
// 		} else if word[i] == "(bin)" {
// 			data, err := strconv.ParseInt(word[i-1], 2, 64)
// 			if err != nil {
// 				return "Error"
// 			}

// 			word[i-1] = strconv.FormatInt(data, 10)

// 			word = append(word[:i], word[i+1:]...)
// 			i--
// 		} else if word[i] == "(up)" {
// 			word[i-1] = strings.ToUpper(word[i-1])

// 			word = append(word[:i], word[i+1:]...)
// 			i--
// 		} else if word[i] == "(cap)" {

// 			runes := []rune(word[i-1])
// 			runes[0] = unicode.ToUpper(runes[0])
// 			word[i-1] = string(runes)

// 			word = append(word[:i], word[i+1:]...)
// 			i--
// 		}
// 	}

// 	result := strings.Join(word, " ")
// 	return result
// }

// func main(){
// 	fmt.Println(TextFormatter("1E (hex) files and hello (up) world (cap) were added"))
// }

//MAKE BUILT-IN FUNCTION

//make(Type, Length)

// package main // 1. Declare the package first

// import "fmt"

// func main() {
//     // 2. Use make() inside a function
//     mySlice := make([]int, 6)
// 	mySlice[0] = 5
//     fmt.Println(mySlice)
// }

//cleanstr

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main(){
// 	if len(os.Args) != 2 {
// 		fmt.Print("\n")
// 		return
// 	}

// 	text := os.Args[1]

// 	if len(text) == 0 {
// 		fmt.Print("\n")
// 		return
// 	}

// 	word := strings.Fields(text)

// 	result := strings.Join(word, " ")

// 	fmt.Println(result)
// }

//expandstr

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main(){
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	text := os.Args[1]

// 	if len(text) == 0 {
// 		return
// 	}

// 	word := strings.Fields(text)

// 	result := strings.Join(word, "   ")

// 	fmt.Println(result)
// }

//findprevprime

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func FindPrevPrime(nb int) int {
// 	for i := nb; i > 1; i-- {
// 		if isPrime(i) {
// 			return i
// 		}
// 	}

// 	return 0
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}

// 	if n == 2 {
// 		return true
// 	}

// 	if n % 2 == 0 {
// 		return false
// 	}

// 	limit := int(math.Sqrt(float64(n)))

// 	for i := 3; i <= limit; i += 2 {
// 		if n % i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// //ANOTHER WAY TO FIND PRIME

// 	for i := 3; i * i <= n; i+=2 {
// 		if n % i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// 	fmt.Println(isPrime(6768765))
// }

/*

FROMTO

package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	if from <= to {
		for i := from; i <= to; i++ {
			result += fmt.Sprintf("%02d", i)
			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			result += fmt.Sprintf("%02d", i)
			if i != to {
				result += ", "
			}
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

*/

// IsCapitalized

/*
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	word := strings.Fields(s)

	for _, r := range word {
		firstChar := rune(r[0])
		if unicode.IsLower(firstChar) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

*/

/*
package main

import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
    negative = true
    n = -n
	}

	result := ""
	for n > 0 {
    digit := n % 10
    result = string(rune('0'+digit)) + result
    n = n / 10
	}

	if negative {
    result = "-" + result
	}

	return result
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
*/
/*
package main

import "fmt"

func PrintMemory(arr [10]byte) {
    // First loop - print hex values 4 per line
    for i, b := range arr {
		if i % 4 != 0 {
    	fmt.Printf(" ")
		}

        fmt.Printf("%02x", b)

        // if end of group of 4, print newline
        if i % 4 == 3 {
            fmt.Println()
        }

    }

	fmt.Println()


    // Second loop - print printable chars or dot
    for _, b := range arr {
        if b >= 32 && b <= 126 {
            fmt.Printf("%c", b)
        } else {
            fmt.Printf(".")
        }
    }
    fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
*/

// THIRDTIMEISACHARM

/*

package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	result := ""

	for i, r := range str {
		if i % 3 == 2 {
			result += string(r)
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
*/

//WEAREUNIQUE

/*

package main

import (
	"fmt"
	"strings"
)

func WeAreUnique(str1 , str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	count := 0

	seen := map[rune]bool{}

	for _, r := range str1 {
		if !strings.ContainsRune(str2, r) && !seen[r] {
			seen[r] = true
			count++
		}
	}

	for _, r := range str2 {
    if !strings.ContainsRune(str1, r) && !seen[r] {
        seen[r] = true
        count++
    }
	}

	return count

}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

*/

/*


package main

import (
	"fmt"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""

	current := rune(s[0])
	count := 1

	for _, r := range s[1:] {
		if current == r {
			count++
		} else if current != r {
			result += fmt.Sprintf("%d%c", count, current)
			count = 1
			current = r
		}
	}

	result += fmt.Sprintf("%d%c", count, current)
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

*/

/*

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	num := os.Args[1]

	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(0)
		return
	}

	sum := 0

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}


func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	for i := 3; i * i <= n; i+=2 {
		if n % i == 0 {
			return false
		}
	}

	return true
}

*/