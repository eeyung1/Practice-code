package converter

import "strings"

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	for _, c := range input {
		if c != '\n' && (c < '0' || c > '9') {
			return ""
		}
	}

	patterns := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    |",
			" ___/",
			"|    ",
			"|___ ",
		},
		'3': {
			" ___ ",
			"    |",
			" __ |",
			"    |",
			"|___|",
		},
		'4': {
			"|   |",
			"|   |",
			"|___|",
			"    |",
			"    |",
		},
		'5': {
			" ___",
			"|    ",
			"|___ ",
			"    |",
			"|___|",
		},
		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		},
		'7': {
			" ___/",
			"   / ",
			"  /  ",
			" /   ",
			"/    ",
		},
		'8': {
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			"|___|",
		},
	}

	lines := strings.Split(input, "\n")
	var output string

	for _, line := range lines {
		rows := make([]string, 5)

		for _, c := range line {
			pattern := patterns[c]
			for i := 0; i < 5; i++ {
				rows[i] += pattern[i]
			}
		}

		output += strings.Join(rows, "\n") + "\n"
	}

	return output
}