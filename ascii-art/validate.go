package main

import "errors"

func ValidateInput(s string) (rune, error) {
	for _, r := range s {
		if r < 32 || r > 126 {
			return r, errors.New("rune is out of range")
		}
	}

	return 0, nil
}
