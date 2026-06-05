package ascii

import (
	"os"
	"strings"
)

func GenerateAscii(text, banner string) (string, error) {


	filepath := "banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)

	if err != nil {

		return "", err
	}

	
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	
	
	lines := strings.Split(content, "\n")
	
	characters := strings.ReplaceAll(text, "\r\n", "\n")

	
	if strings.ContainsAny(text, `\n`) {
		characters = strings.ReplaceAll(text, `\n`, "\n")
	}
	parts := strings.Split(characters, "\n")

	result := ""

	for _, part := range parts {

		for i := 0; i < 8; i++ {

			for _, ch := range part {

				index := int(ch-32)*9 + i

				result += lines[index]
			}

			result += "\n"
		}

		result += "\n"
	}

	return result, nil
}
