package main

import "fmt"

func validateInput(banner map[rune][]string) error {
	if len(banner) != 95 {
		return fmt.Errorf("banner has %d enteries,expected 95", len(banner))
	}
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}

	for r := rune(32); r <= 126; r++ {
		lines, ok := banner[r]
		if !ok {
			return fmt.Errorf("missing character: '%c' (Ascii %d)", r, r)
		}
		if len(lines) != 8 {
			return fmt.Errorf("character '%c' has %d lines, expected 8", r, len(lines))
		}
	}
	return nil
}
