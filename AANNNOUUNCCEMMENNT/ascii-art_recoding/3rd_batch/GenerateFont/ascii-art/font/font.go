package font

import "fmt"

func pad(s string) string {
    for len(s) < 8 {
        s += " "
    }

    return s[:8]
}

func GenerateFont() map[rune][]string {
    font := make(map[rune][]string)

    for c := rune(32); c <= rune(126); c++ {
        font[c] = generate(c)
    }

    return font
}

func generate(c rune) []string {
    result := make([]string, 8)

    // Space must be pure blank
    if c == 32 {
        for i := range result {
            result[i] = "        "
        }
        return result
    }

    // Encode ASCII value into a unique 8-row pattern
    val := int(c)
    for i := 0; i < 8; i++ {
        // Each row is determined by one bit of the ASCII value
        if (val>>i)&1 == 1 {
            result[i] = pad(fmt.Sprintf("*%d*%d*%d*", i, val%8, i))
        } else {
            result[i] = pad(fmt.Sprintf(".%d.%d.%d.", i, val%8, i))
        }
    }

    return result
}

/*package font

func pad(s string) string {
	for len(s) < 8 {
		s += " "
	}
	return s[:8]
}

func lines(l ...string) []string {
	result := make([]string, 8)
	for i := 0; i < 8; i++ {
		if i < len(l) {
			result[i] = pad(l[i])
		} else {
			result[i] = "        "
		}
	}
	return result
}

func spacePattern() []string {
	return lines(
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	)
}

func vowelPattern(c rune) []string {
	o := int(c) % 4
	mid := "        "
	if o%2 == 0 {
		mid = "*  **  *"
	} else {
		mid = "* *  * *"
	}
	return lines(
		"  ****  ",
		" *    * ",
		"*      *",
		mid,
		"********",
		"*      *",
		" *    * ",
		"  ****  ",
	)
}

func consonantPattern(c rune) []string {
	o := int(c) % 5
	bar := "*      *"
	cross := "****    "
	switch o {
	case 0:
		cross = "*****   "
	case 1:
		cross = " ****   "
	case 2:
		cross = "******  "
	case 3:
		cross = " *****  "
	case 4:
		cross = "******* "
	}
	return lines(
		"*******.",
		bar,
		bar,
		cross,
		bar,
		bar,
		bar,
		"*******.",
	)
}

func lowerPattern(c rune) []string {
	o := int(c) % 4
	ascenders := map[rune]bool{
		'b': true, 'd': true, 'f': true,
		'h': true, 'k': true, 'l': true, 't': true,
	}
	descenders := map[rune]bool{
		'g': true, 'j': true, 'p': true,
		'q': true, 'y': true,
	}

	top := "        "
	bottom := "        "

	if ascenders[c] {
		top = "  *     "
	}
	if descenders[c] {
		bottom = " *      "
	}

	mid := "        "
	switch o {
	case 0:
		mid = " ****   "
	case 1:
		mid = "  ***   "
	case 2:
		mid = " *****  "
	case 3:
		mid = "  ****  "
	}

	return lines(
		top,
		top,
		"        ",
		" *****  ",
		"*     * ",
		mid,
		"*       ",
		bottom,
	)
}

func digitPattern(c rune) []string {
	o := int(c-'0') // 0 through 9 exactly
	inner := "        "
	switch o % 3 {
	case 0:
		inner = "|      |"
	case 1:
		inner = "|  **  |"
	case 2:
		inner = "| *  * |"
	}
	return lines(
		" ______ ",
		"|      |",
		inner,
		inner,
		"|  "+string(rune('0'+o))+"   |",
		inner,
		"|      |",
		"|______|",
	)
}

func symbolPattern(c rune) []string {
	o := int(c) % 6
	row := pad(string(c) + string(c) + string(c))
	empty := "        "
	result := make([]string, 8)
	for i := range result {
		result[i] = empty
	}
	result[o] = row
	result[o+1] = row
	if o+2 < 8 {
		result[o+2] = row
	}
	return result
}

func isUpperVowel(c rune) bool {
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func isLowerVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for c := rune(32); c <= rune(126); c++ {
		switch {
		case c == ' ':
			font[c] = spacePattern()
		case c >= '0' && c <= '9':
			font[c] = digitPattern(c)
		case c >= 'A' && c <= 'Z':
			if isUpperVowel(c) {
				font[c] = vowelPattern(c)
			} else {
				font[c] = consonantPattern(c)
			}
		case c >= 'a' && c <= 'z':
			font[c] = lowerPattern(c)
		default:
			font[c] = symbolPattern(c)
		}
	}

	return font
}

*/