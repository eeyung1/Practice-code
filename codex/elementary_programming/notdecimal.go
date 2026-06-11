package main

import "fmt"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Step 1: Validate - only digits, one optional leading '-', one optional '.'
	dotCount := 0
	for i, ch := range dec {
		if ch == '-' && i == 0 {
			continue
		}
		if ch == '.' {
			dotCount++
			if dotCount > 1 {
				return dec + "\n"
			}
			continue
		}
		if ch < '0' || ch > '9' {
			return dec + "\n"
		}
	}

	// Step 2: find the dot's position
	dotIndex := -1
	for i, ch := range dec {
		if ch == '.' {
			dotIndex = i
			break
		}
	}
	if dotIndex == -1 {
		return dec + "\n"
	}

	// Step 3: check if everything after the dot is zero
	allZero := true
	for _, ch := range dec[dotIndex+1:] {
		if ch != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return dec[:dotIndex] + "\n"
	}

	// Step 4: build combined string without the dot
	combined := ""
	for _, ch := range dec {
		if ch != '.' {
			combined += string(ch)
		}
	}

	// Step 5: strip leading zeros, preserving sign
	negative := false
	start := 0
	if combined[0] == '-' {
		negative = true
		start = 1
	}
	for start < len(combined)-1 && combined[start] == '0' {
		start++
	}
	result := combined[start:]
	if negative {
		result = "-" + result
	}

	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}