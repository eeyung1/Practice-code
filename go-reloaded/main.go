package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. Validate arguments
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . input.txt output.txt")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    // 2. Read input file
    data, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    // 3. Transform the text
    result := transform(string(data))

    // 4. Write to output file
    err = os.WriteFile(outputFile, []byte(result), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        os.Exit(1)
    }
}



// every transformation follows the same rhythm:
// 1. Build a regex to FIND the pattern
// 2. Extract the PARTS you need
// 3. TRANSFORM those parts
// 4. RETURN the new string