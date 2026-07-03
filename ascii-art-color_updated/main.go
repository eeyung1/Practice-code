package main

import (
	ascii "color/ascii_art"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	arg := os.Args[1:]

	if len(arg) < 1 || len(arg) > 4 {
		Usage()
		return
	}

	var (
		text      string
		banner    = "standard"
		color     = ""
		substring = ""
	)

	switch len(arg) {
	case 1:

		text = arg[0]

	case 2:
		if !strings.HasPrefix(arg[0], "--color=") {
			text = arg[0]

			if IsBanner(arg[1]) {
				banner = arg[1]
			} else {
				log.Fatal("Invalid banner:", banner)
			}

		} else {

			ansi, ok := GetANSI(arg[0])
			if !ok {
				Usage()
				return
			}

			color = ansi
			text = arg[1]
		}

	case 3:
		if !strings.HasPrefix(arg[0], "--color=") {
			Usage()
			return
		}

		ansi, ok := GetANSI(arg[0])

		if !ok {
			Usage()
			return
		}

		color = ansi

		if IsBanner(arg[2]) {

			text = arg[1]
			banner = arg[2]

		} else {

			substring = arg[1]
			text = arg[2]

		}

	case 4:

		if !strings.HasPrefix(arg[0], "--color=") {
			Usage()
			return
		}

		ansi, ok := GetANSI(arg[0])

		if !ok {
			Usage()
			return
		}

		color = ansi

		substring = arg[1]
		text = arg[2]

		if IsBanner(arg[3]) {
			banner = arg[3]
		} else {
			log.Fatal("Invalid banner:", banner)
		}

	default:
		Usage()
		return
	}

	bannerFile, err := ascii.LoadBanner("banners/" + banner)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(
		ascii.GenerateArt(
			text,
			bannerFile,
			color,
			substring,
		),
	)

}

func Usage() {

	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")

}

func ParseColor(flag string) (string, bool) {

	if !strings.HasPrefix(flag, "--color=") {
		return "", false
	}

	return strings.TrimPrefix(flag, "--color="), true
}

func IsBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}

func GetANSI(flag string) (string, bool) {
    name, ok := ParseColor(flag)
    if !ok {
        return "", false
    }

    return ascii.GetColor(name)
}