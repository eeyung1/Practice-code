package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]Room, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rooms []Room

	roomLookup := make(map[string]Room)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		x, err := strconv.Atoi(parts[1])

		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		room := Room{
			Name: parts[0],
			X:    x,
			Y:    y,
		}

		if _, exists := roomLookup[room.Name]; exists {
			return nil, fmt.Errorf("duplicate room name: %s", room.Name)
		}

		rooms = append(rooms, room)
		roomLookup[room.Name] = room
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
