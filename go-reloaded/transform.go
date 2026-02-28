package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func applyHex(text string) string {
	re := regexp.MustCompile(`([0-9a-fA-F]+)\s+\(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", num)
	})
}

func applyBin(text string) string {
	re := regexp.MustCompile(`([0-1]+)\s+\(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 2, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", num)
	})
}

func applyCases(text string) string {
	re := regexp.MustCompile(`(?i)\((up|low|cap),?\s*(\d+)?\)`)

	for re.MatchString(text) {
		match := re.FindStringSubmatchIndex(text)
		if match == nil {
			break
		}

		// extract action and count
		action := strings.ToLower(text[match[2]:match[3]])
		count := 1
		if match[4] != -1 {
			count, _ = strconv.Atoi(text[match[4]:match[5]])
		}

		// get the text before the tag
		before := text[:match[0]]
		after := text[match[1]:]

		// split before into words and transform the last `count` words
		words := strings.Fields(before)
		for i := len(words) - count; i < len(words); i++ {
			if i < 0 {
				continue
			}
			switch action {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				words[i] = capitalize(words[i])
			}
		}

		text = strings.Join(words, " ") + after
	}
	return text
}


func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}


func applyArticle(text string) string {
	re := regexp.MustCompile(`(?i)\b(a)\b(\s+)([aeiouAEIOUhH]\w*)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		// preserve the original case of "a" or "A"
		article := parts[1]
		space := parts[2]
		nextWord := parts[3]
		if article == "A" {
			return "An" + space + nextWord
		}
		return "an" + space + nextWord
	})
}

func applyPunctuation(text string) string {
	// step 1: remove spaces before punctuation
	re1 := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = re1.ReplaceAllString(text, "$1")

	// step 2: ensure space after punctuation
	re2 := regexp.MustCompile(`([.,!?:;]+)(\s*)(\w)`)
	text = re2.ReplaceAllString(text, "$1 $3")

	return text
}

func applyQuotes(text string) string {
	re := regexp.MustCompile(`'\s+(.+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

func transform(text string) string {
	text = applyHex(text)
	text = applyBin(text)
	text = applyCases(text)
	text = applyArticle(text)
	text = applyPunctuation(text)
	text = applyQuotes(text)
	return text
}


