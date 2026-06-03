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

	result := ""

	for i := 0; i < 8; i++ {

		for _, ch := range text {

			index := int(ch-32)*9 + i

			result += lines[index]
		}

		result += "\n"
	}

	return result, nil
}