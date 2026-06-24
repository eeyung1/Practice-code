package main

import "fmt"

func CaesarCipher(text string, shift int) string {
    answer := []rune{}
    
    for _, ch := range text {
        if ch >= 'a' && ch <= 'z' {
            // Shift lowercase
            value := ch - 'a'
            result := value + rune(shift)
            finalResult := result % 26
            char := finalResult + 'a'
            answer = append(answer, char)
            
        } else if ch >= 'A' && ch <= 'Z' {
            // Shift uppercase
            value := ch - 'A'
            result := value + rune(shift)
            finalResult := result % 26
            char := finalResult + 'A'
            answer = append(answer, char)
            
        } else {
            // Preserve everything else
            answer = append(answer, ch)
        }
    }
    
    return string(answer)
}

func main() {
    text := "Hello, World!"
    shift := 3
    
    encrypted := CaesarCipher(text, shift)
    fmt.Printf("Original:  %s\n", text)
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Try to decrypt
    decrypted := CaesarCipher(encrypted, -shift)
    fmt.Printf("Decrypted: %s\n", decrypted)
}