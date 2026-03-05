package main

import (
	"fmt"
	"strings"
	"strconv"
)


// func main() {
// text := "It has been 10 (bin) years"

// textsplit := strings.Fields(text)

// for i := 0; i < len(textsplit); i++ {
// 	if textsplit[i] == "(bin)" {
// 		editor := textsplit[i-1]

// 		num, err := strconv.ParseInt(editor, 2, 64)
// 		if err != nil {
// 			fmt.Println("Error reading file")
// 			return
// 		}
			
// 		textsplit[i-1] = fmt.Sprintf("%d", num)
// 		textsplit = append(textsplit[:i], textsplit[i+1:]...)

// 	}
// }

// result := strings.Join(textsplit, " ")
// fmt.Println(result)

// }

func main() {
text := "1E (hex) plus FF (hex) equals" 

textsplit := strings.Fields(text)

for i := 0; i < len(textsplit); i++ {
	if textsplit[i] == "(hex)" {
		editor := textsplit[i-1]

		num, err := strconv.ParseInt(editor, 16, 64)
		if err != nil {
			fmt.Println("Error reading file")
			return
		}
			
		textsplit[i-1] = fmt.Sprintf("%d", num)
		textsplit = append(textsplit[:i], textsplit[i+1:]...)

	}
}

result := strings.Join(textsplit, " ")
fmt.Println(result)

}


