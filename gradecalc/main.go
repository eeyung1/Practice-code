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

func processStudent(row string) {
	// step 1: split by comma
	values := strings.Split(row, ",")
	name := values[0]

	// step 2: convert grades to float64
	grade1, _ := strconv.ParseFloat(values[1], 64)
	grade2, _ := strconv.ParseFloat(values[2], 64)
	grade3, _ := strconv.ParseFloat(values[3], 64)

	// step 3: calculate average
	avg := (grade1 + grade2 + grade3) / 3

	// step 4: find highest
	highest := grade1
	if grade2 > highest {
		highest = grade2
	}
	if grade3 > highest {
		highest = grade3
	}

	// step 5: find lowest
	lowest := grade1
	if grade2 < lowest {
		lowest = grade2
	}
	if grade3 < lowest {
		lowest = grade3
	}

	fmt.Printf("%-10s avg: %.2f  highest: %.0f  lowest: %.0f\n", name, avg, highest, lowest)
}