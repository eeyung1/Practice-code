package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Generate(text string, bannerType string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("text cannot be empty")
	}

	validBanners := map[string]bool{"standard": true, "shadow":true, "thinkertoy": true}
	if !validBanners[bannerType] {
		return "", fmt.Errorf("invalid banner type: %s", bannerType)
	}

	filepath := "banners/" + bannerType + ".txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("banner file not found: %s", bannerType)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	var result strings.Builder

	for i := 0; i < 8; i++ {
		for _, ch := range text {
			if ch < 32 || ch > 126 {
				return "", fmt.Errorf("unsupported character: %c", ch)
			}

			index := int(ch-32)*9 + 1 + i

			if index >= len(lines) {
				return "", fmt.Errorf("banner file format error")
			}

			result.WriteString(lines[index])
		}

		result.WriteString("\n")
	}

	return result.String(), nil
}

