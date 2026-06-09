package main

import (
	"fmt"
	"os"
	"strings"
)

func artbuilder(text, bannerfile string) string {
	fmt.Println(text)
	good := []rune(text)
	fmt.Println(good)

	text = strings.ReplaceAll(text, "\r\n", "\n")

	banner, err := os.ReadFile("banners/" + bannerfile)

	if err != nil {
		return fmt.Sprintf("error reading file: %v", err)
	}

	fileline := strings.Split(strings.ReplaceAll(string(banner), "\r\n", "\n"), "\n")

	var build strings.Builder

	splitinput := strings.Split(text, "\n")

	for _, word := range splitinput {
		for row := 0; row < 8; row++ {
			for _, ch := range word {

				if ch < 32 || ch > 126 {
					continue
				}

				start := int(ch-32) * 9

				end := start + 8

				build.WriteString(fileline[start:end][row])
			}
			build.WriteString("\n")
		}
	}
	return build.String()
}
