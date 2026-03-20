// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func BaseConv(num string) (int64, error, string) {
// 	word := strings.Fields(num)
// 	for i := 0; i < len(word); i++ {
// 		if word[i] == "(hex)" {
// 			data, err := strconv.ParseInt(word[i-1], 16, 64)
// 			if err != nil {
// 				return 0, err, "logic"
// 			}

// 			return data, nil, "logic"

// 		} else if word[i] == "(bin)" {
// 			data, err := strconv.ParseInt(word[i-1], 2, 64)
// 			if err != nil {
// 				return 0, err, "logic"
// 			}

// 			return data, nil, "logic"
// 		}
// 	}

// 	return 0, nil, ""
// }

// func main() {
// 	fmt.Println(BaseConv("FF (hex)"))
// 	fmt.Println(BaseConv("101 (bin)"))
// }

package main

import (
	"fmt"
)

func BinToDecimal(bin string) (string, error) {
	data, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d", data) //result := strconv.FormatInt(data, 10)
	return result, nil
}

func main() {
	fmt.Println(BinToDecimal("11"))
	fmt.Println(BinToDecimal("1001"))
}
