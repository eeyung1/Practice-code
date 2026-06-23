package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	firstArg := os.Args[1]

	if len(firstArg) >= 2 && firstArg[0] == '-' && firstArg[1] == 'h' {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	for _, arg := range os.Args[1:] {
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

		if arg[0] == '-' && arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

	}

	options := 0

	for _, arg := range os.Args[1:] {
		if arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		if len(arg) == 1 {
			fmt.Println("Invalid Option")
			return
		}

		for i := 1; i < len(arg); i++ {
			ch := arg[i]

			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			pos := int(ch - 'a')

			options = options | (1 << pos)
			fmt.Println(options)
			
		}
	}

	fmt.Printf("%08b %08b %08b %08b\n",
		byte(options>>24),
		byte(options>>16),
		byte(options>>8),
		byte(options&0xFF))
}
