package main

import (
    "fmt"
    "time"
)

func serveTable(table int) {
    fmt.Printf("Starting table %d\n", table)
    time.Sleep(2 * time.Second)
    fmt.Printf("Finished table %d\n", table)
}

func main() {
    serveTable(1)
    serveTable(2)
    serveTable(3)

    time.Sleep(3 * time.Second)
    fmt.Println("All done")
}