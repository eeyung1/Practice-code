package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`hello`)
	fmt.Println(re.MatchString("hello world"))
	fmt.Println(re.MatchString("goodbye world"))
}

// package main

// import (
//     "fmt"
//     "regexp"
// )

// func main() {
//     re := regexp.MustCompile(`([0-9a-fA-F]+)\s+\(hex\)`)
//     result := re.FindStringSubmatch("1E (hex) files added")
//     fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	re := regexp.MustCompile(`([0-9a-fA-F]+)\s+\(hex\)`)
// 	result := re.ReplaceAllStringFunc("1E (hex) files and FF (hex) added", func(match string) string {
// 		fmt.Println("match:", match)
// 		return "REPLACED"
// 	})

// 	fmt.Println(result)
// }
