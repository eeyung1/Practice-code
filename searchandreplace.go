package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		str, old, new_sub := os.Args[1], os.Args[2], os.Args[3]

		if len(old) == 1 && len(new_sub) == 1 {
			if strings.Contains(str, old) {
				str = strings.ReplaceAll(str, old, new_sub)
			}

			fmt.Println(str)
		}
	}
}
