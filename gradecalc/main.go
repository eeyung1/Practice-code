package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// step 1: confirm arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . students.csv")
		os.Exit(1)
	}

	// step 2: read the file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	text := strings.TrimRight(string(data), "\n")

	// step 3: split into rows
	rows := strings.Split(text, "\n")

	// step 4: print each row to verify
	// replace the print loop with this
	for _, row := range rows {
		processStudent(row)
	}
}

// handles full row — used by main()
func processStudent(row string) float64 {
    values := strings.Split(row, ",")
    name := values[0]
    grade1, _ := strconv.ParseFloat(values[1], 64)
    grade2, _ := strconv.ParseFloat(values[2], 64)
    grade3, _ := strconv.ParseFloat(values[3], 64)

    grades := []float64{grade1, grade2, grade3}
    avg, highest, lowest := calculateStats(grades)  // ← call calculateStats

    fmt.Printf("%-10s avg: %.2f  highest: %.0f  lowest: %.0f\n", name, avg, highest, lowest)
    return avg
}

// pure math logic — easy to test
func calculateStats(grades []float64) (float64, float64, float64) {
    highest := grades[0]
    lowest := grades[0]
    total := 0.0

    for _, g := range grades {
        total += g
        if g > highest {
            highest = g
        }
        if g < lowest {
            lowest = g
        }
    }

    avg := total / float64(len(grades))
    return avg, highest, lowest
}