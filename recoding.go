// RECODING QUESTIONS

// 1. convertToDecimal

/*
package main

import (
    "fmt"
    "strconv"
)

func convToDecimal(num string, base int) (int64, error) {
    decimal, err := strconv.ParseInt(num, base, 64)
    if err != nil {
        return 0, err
    }

    return decimal, nil
}

func main(){
    fmt.Println(convToDecimal("1E", 16))
    fmt.Println(convToDecimal("1010", 2))
	    fmt.Println(convToDecimal("77", 8))
}

*/

// 2. Capitalize a word
/*
package main

import (
	"fmt"
	"strings"
)

func Capitalize(sentence string) string {
	word := strings.Fields(sentence)

	for i := 0; i < len(word); i++ {
		word[i] = strings.Title(word[i])
	}

	return strings.Join(word, " ")
}

func main() {
	fmt.Println(Capitalize("hello world from golang"))
}

*/

// 3. Last two words to upper case

/*
package main

import (
    "fmt"
    "strings"
)
func LastTwoToUpper(s []string, n int) []string {
    for i := n; i < len(s); i++ {
        s[i] = strings.ToUpper(s[i])
    }
    return s
}

func main () {
    fmt.Println(LastTwoToUpper([]string{"This", "is", "so", "exciting"}, 2))
}

*/

// 4. joinWithPunctuation

/*

package main

import (
	"fmt"
	"strings"
)

func joinWithPunctuation(tokens []string) string {
    result := strings.Join(tokens, "")

    // Replace the comma with a comma and a space
    return strings.ReplaceAll(result, ",", ", ")
}


func main(){	// Removes the space after the opening single quote

	fmt.Println(joinWithPunctuation([]string{"hello", ",", "world", "!"}))
}


*/

// 5. IsPunctuation

/*
package main

import (
	"fmt"
)

func isPunctuation(s string) bool {
	for _, r := range s {
		if r == ',' || r == '!' {
			return true
		}
	}
	return false
}

func main(){
	fmt.Println(isPunctuation(","))
	fmt.Println(isPunctuation("!"))
	fmt.Println(isPunctuation("x"))
}
*/

// 6. func aOrAn

/*

package main

import (
	"fmt"
)

func aOrAn(nextWord string) string {
	if nextWord == "honest" {
		return "an"
	}

	for _, r := range nextWord {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' {
			return "an"
		} else {
			return "a"
		}
	}

	return ""
}

func main() {
	fmt.Println(aOrAn("good"))
	fmt.Println(aOrAn("honest"))
	fmt.Println(aOrAn("apple"))
	fmt.Println(aOrAn("yam"))
	fmt.Println(aOrAn("horse"))

}

*/

// 7. fixArticles

/*
package main

import (
	"fmt"
	"strings"
)

func fixArticles(text string) string {

	text = strings.ReplaceAll(text, "A", "An")
	text = strings.ReplaceAll(text, "An book.", "A book.")

	return text
}

func main(){
	fmt.Println(fixArticles("There it was. A amazing rock. A honest man. A book."))
}

*/

// 8. fixSingleQuotes

/*
package main

import (
	"fmt"
	"strings"
)

func fixSingleQuotes(text string) string {
	text = strings.Trim(text, "'")

	text = strings.TrimSpace(text)

	return "'" + text + "'"
}

func main() {
	input := "' hello world '"

	good := "' awesome '"

	great := fixSingleQuotes(good)

	result := fixSingleQuotes(input)

	fmt.Printf("%q\n", result)
	fmt.Printf("%q\n", great)
}

*/
