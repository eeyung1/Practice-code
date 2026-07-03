package ascii

import "strings"

const Reset = "\033[0m"

var Colors = map[string]string{
	"black":  "\033[30m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
}

func GetColor(name string) (string, bool) {
	color, ok := Colors[strings.ToLower(name)]
	return color, ok
}

func BuildHighlight(text, sub string) []bool {
	highlight := make([]bool, len(text))

	if sub == "" {
		for i := range highlight {
			highlight[i] = true
		}
		return highlight
	}

	start := 0

	for {
		index := strings.Index(text[start:], sub)

		if index == -1 {
			break
		}

		index += start

		for i := 0; i < len(sub); i++ {
			highlight[index+i] = true
		}

		start = index + 1
	}

	return highlight
}