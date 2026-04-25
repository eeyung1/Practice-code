package main

import (
	"fmt"
)

func main() {
	m := map[rune][]string{}
	m['A'] = []string{
		"  A  ",
		" A A ",
		"AAAAA",
		"A   A",
		"A   A",
	}
	m['B'] = []string{
		"BBBB ",
		"B   B",
		"BBBBB",
		"B   B",
		"BBBB ",
	}
	m['C'] = []string{
		" CCCC",
		"C    ",
		"C    ",
		"C    ",
		" CCCC",
	}
	m['D'] = []string{
		"DDDD ",
		"D   D",
		"D   D",
		"D   D",
		"DDDD",
	}
	m['E'] = []string{
		"EEEEE",
		"E    ",
		"EEEEE",
		"E    ",
		"EEEEE",
	}

	input := "BAD"

	for i := 0; i < len(m['A']); i++ {
		for _, ch := range input {
			fmt.Print(m[ch][i], "  ")
		}
		fmt.Println()
	}
}


// Read a banner file and print a character

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")

	input := 'A'
	start := (int(input)-32)*9 + 1
	end := start + 8
	res := append(lines[start:end])
	for _, line := range res {
		fmt.Println(line)
	}
}

//-----------OR--------------

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")
	input := 'A'
	read(input, lines)
}

func read(input rune, lines []string) {
	start := (int(input)-32)*9 + 1
	end := start + 8
	res := lines[start:end]
	for _, line := range res {
		fmt.Println(line)
	}
}

//---------------OR--------------

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")

	input := '%'

	start := (int(input)-32)*9 + 1
	end := start + 8

	res := lines[start:end]
	fmt.Println(strings.Join(res, "\n"))
}





// Read a banner file and print a string
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")
	input := "HELLO&"
	printAscii(input, lines)
}

func getCharArt(lines []string, ch rune) []string {
	start := (int(ch) - 32) * 9
	return lines[start : start+8]
}
func printAscii(input string, banner []string) {
	for i := 0; i < 8; i++ {
		for _, ch := range input {
			art := getCharArt(banner, ch)
			fmt.Print(art[i])
		}
		fmt.Println()
	}
}

//-------------OR---------------

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")
	input := "HELLO&"
	printAscii(input, lines)
}

func printAscii(input string, banner []string) {
	for i := 0; i < 8; i++ {
		for _, ch := range input {
			start := (int(ch) - 32) * 9
			art := banner[start : start+8]
			fmt.Print(art[i])
		}
		fmt.Println()
	}
}

//-------------OR-------------

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")

	input := "ABC"
	res := make([]string, 8)

	for _, char := range input {
		start := (int(char)-32)*9 + 1

		for i := 0; i < 8; i++ {
			res[i] += lines[start+i]
		}
	}

	fmt.Println(strings.Join(res, "\n"))
}




// Single character - memory management

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(banner), "\n")

	var result strings.Builder
	for i := 0; i < 8; i++ {
		index := ('A'-32)*9 + 1 + i

		result.WriteString(lines[index] + "\n")
	}

	fmt.Print(result.String())
}
