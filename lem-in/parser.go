package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) (*ParsedData, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	parsed := &ParsedData{}

	scanner := bufio.NewScanner(file)

	roomLookup := make(map[string]Room)

	nextRoomType := ""

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") && line != "##start" && line != "##end" {
			continue
		}

		if line == "##start" {
			nextRoomType = "start"
			continue
		}

		if line == "##end" {
			nextRoomType = "end"
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 1 {
			ants, err := strconv.Atoi(parts[0])

			if err != nil {
				return nil, fmt.Errorf("invalid ant count: %s", parts[0])
			}

			parsed.Ants = ants
			continue
		}

		if len(parts) != 3 {
			continue
		}

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

		if nextRoomType == "start" {
			room.IsStart = true
			nextRoomType = ""
		} else if nextRoomType == "end" {
			room.IsEnd = true
			nextRoomType = ""
		}

		if _, exists := roomLookup[room.Name]; exists {
			return nil, fmt.Errorf("duplicate room name: %s", room.Name)
		}

		parsed.Rooms = append(parsed.Rooms, room)
		roomLookup[room.Name] = room
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parsed, nil
}
