package main

import (
	"fmt"
	"os"
	"ascii-art/factors"

)

const defaultBanner = "standard"

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <text> [banner]")
		fmt.Fprintln(os.Stderr, "Available banners: standard, shadow, thinkertoy")
		os.Exit(1)
	}

	input := args[0]

	bannerName := "banners/" + defaultBanner
	if len(args) == 2 {
		bannerName = "../banners/" + args[1]
	}

	bannerFile := bannerName + ".txt"

	banner, err := factors.LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not load banner %q: %v\n", bannerName, err)
		os.Exit(1)
	}

	result := factors.GenerateArt(input, banner)
	fmt.Print(result)
}