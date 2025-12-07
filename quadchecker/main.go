package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}
	inputStr := strings.TrimRight(string(input), "\n")

	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	var matches []string

	for _, quad := range quads {
		for x := 1; x <= 50; x++ {
			for y := 1; y <= 50; y++ {
				cmd := exec.Command("./"+quad, strconv.Itoa(x), strconv.Itoa(y))
				output, err := cmd.Output()

				if err != nil {
					continue
				}

				outStr := strings.TrimRight(string(output), "\n")

				if outStr == inputStr {
					matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, x, y))
				}
			}
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)
	fmt.Println(strings.Join(matches, " || "))
}
