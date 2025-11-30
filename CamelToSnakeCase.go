package main

import "fmt" // Imported as requested

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	res := ""
	prevUpper := false
	for i, r := range s {
		isLower, isUpper := r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z'
		if !isLower && !isUpper {
			return s
		}
		if isUpper && prevUpper {
			return s
		}

		if isUpper && i != 0 {
			res += "_"
		}

		res += string(r)
		prevUpper = isUpper
	}
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}
	return res
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
