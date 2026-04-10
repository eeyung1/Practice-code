package main

import (
	"fmt"
	"unicode"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}

	if unicode.IsUpper(rune(s[len(s)-1])) {
		return s
	}

	prevUpper := false
	var result []rune

	for i, r := range s {
		if (!unicode.IsUpper(r) && !unicode.IsLower(r)) || (unicode.IsUpper(r) && prevUpper) {
			return s
		}

		if unicode.IsUpper(r) && i != 0 {
			result = append(result, '_')
		}

		result = append(result, r)
		prevUpper = unicode.IsUpper(r)
	}

	return string(result)
}

func SnakeToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	capitalizeNext := false
	var result []rune

	for i, r := range s {
		if r != '_' && (r < 'a' || r > 'z') {
			return s
		}

		if r == '_' {
			if i == 0 || i == len(s)-1 || s[i-1] == '_' {
				return s
			}

			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			result = append(result, r-32)
			capitalizeNext = false
		} else {
			result = append(result, r)
		}
	}

	return string(result)

}

func CamelToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	if unicode.IsUpper(rune(s[len(s)-1])) {
		return s
	}

	prevUpper := false
	var result []rune

	for i, r := range s {
		if !unicode.IsLetter(r) || (unicode.IsUpper(r) && prevUpper) {
			return s
		}

		if unicode.IsUpper(r) && i != 0 {
			result = append(result, '-')
		}

		result = append(result, unicode.ToLower(r))
		prevUpper = unicode.IsUpper(r)

	}

	return string(result)
}

func ConvertCase(s string, mode string) string {
	if s == "" && mode == "" {
		return s
	}

	switch mode {
	case "CamelToSnakeCase":
		return CamelToSnakeCase(s)
	case "SnakeToCamelCase":
		return SnakeToCamelCase(s)
	case "CamelToKebabCase":
		return CamelToKebabCase(s)
	default:
		return "Invalid mode"
	}
}

func main() {
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToKebabCase("helloWorld"))
	fmt.Println(CamelToKebabCase("camelCaseTest"))
	fmt.Println(ConvertCase("snakeCase", "CamelToKebabCase"))
}
