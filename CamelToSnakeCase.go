package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	res := ""
	prevUpper := false

	for i, r := range s {
		isLower := r >= 'a' && r <= 'z'
		isUpper := r >= 'A' && r <= 'Z'

		if !isLower && !isUpper || isUpper && prevUpper {
			return s
		}

		if isUpper && i != 0 {
			res += "_"
		}

		res += string(r)
		isUpper = prevUpper

		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			return s
		}
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
