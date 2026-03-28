package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("> ")

    name, _ := reader.ReadString('\n')

    parts := strings.Split(name, " ")


    for i := 0; i < len(parts); i++ {
        parts[i] = strings.Title(parts[i])
    }

    fmt.Println(parts)
    fmt.Println("So this is real ", name)

    
}