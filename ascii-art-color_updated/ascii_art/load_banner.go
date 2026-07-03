package ascii

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(file string) (map[rune][]string, error) {
	banner := map[rune][]string{}

	if !strings.HasSuffix(file, ".txt") {
		file += ".txt"
	}

	data, err := os.ReadFile(file)

	if err != nil {
		return nil, errors.New("Banner file does not exist")
	}

	bannerFile := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(bannerFile) != 856 {
		return nil, errors.New("Invalid banner file")
	}

	const charHeight = 8
	start := 0

	for r := ' '; r <= '~'; r++ {
		
		for start < len(bannerFile) && bannerFile[start] == "" {
			start++
		}

		banner[r] = bannerFile[start : start + charHeight]
		start += charHeight
	}

	return banner, nil
}
