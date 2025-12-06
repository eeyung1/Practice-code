package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func genQuad(w, h int, topLeft, topRight, bottomLeft, bottomRight, edge, middle string) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	lines := make([]string, h)
	for r := 0; r < h; r++ {
		switch r {
		case 0:
			if w == 1 {
				lines[r] = topLeft
			} else {
				lines[r] = topLeft + strings.Repeat(edge, w-2) + topRight
			}
		case h - 1:
			if w == 1 {
				lines[r] = bottomLeft
			} else {
				lines[r] = bottomLeft + strings.Repeat(edge, w-2) + bottomRight
			}
		default:
			if w == 1 {
				lines[r] = middle
			} else {
				lines[r] = middle + strings.Repeat(" ", w-2) + middle
			}
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	raw, _ := io.ReadAll(os.Stdin)
	input := strings.TrimRight(string(raw), "\n")
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)
	for _, L := range lines {
		if len(L) != w {
			fmt.Println("Not a quad function")
			return
		}
	}

	quads := map[string][6]string{
		"quadA": {"o", "o", "o", "o", "-", "|"},
		"quadB": {"/", "\\", "\\", "/", "*", "*"},
		"quadC": {"A", "A", "C", "C", "B", "B"},
		"quadD": {"A", "C", "A", "C", "B", "B"},
		"quadE": {"A", "C", "C", "A", "B", "B"},
	}

	matches := []string{}
	for name, chars := range quads {
		if input == genQuad(w, h, chars[0], chars[1], chars[2], chars[3], chars[4], chars[5]) {
			matches = append(matches, name)
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	sort.Strings(matches)
	out := []string{}
	for _, m := range matches {
		out = append(out, fmt.Sprintf("[%s] [%d] [%d]", m, w, h))
	}
	fmt.Println(strings.Join(out, " || "))
}
