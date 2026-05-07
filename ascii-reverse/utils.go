package main

import "strings"

func isReverseFlag(arg string) bool {
	return strings.HasPrefix(arg, "--reverse=")
}

func getFileName(arg string) string {
	return strings.TrimPrefix(arg, "--reverse=")
}