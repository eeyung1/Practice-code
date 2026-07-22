package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		x, err := strconv.Atoi(parts[1])

		if err != nil {
			return err
		}

		y, err := strconv.Atoi(parts[2])
		if err != nil {
			return err
		}

		room := Room{
			Name: parts[0],
			X:    x,
			Y:    y,
		}

		room.Display()
	}

	return scanner.Err()
}
