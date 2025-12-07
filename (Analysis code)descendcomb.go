package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			z01.PrintRune(rune(i/10) + '0')
			z01.PrintRune(rune(i%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10) + '0')
			z01.PrintRune(rune(j%10) + '0')

		}
	}
}
