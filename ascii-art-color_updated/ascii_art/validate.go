package ascii

import "errors"

func ValidateInput(input string) (rune, error) {

	for _, r := range input {

		if r < ' ' || r > '~' {

			return r, errors.New("Input contains none printable ascii character")

		}
		
	}

	return 0, nil
}
