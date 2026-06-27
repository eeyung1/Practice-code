package main

import "fmt"

func CaesarCipher(text string, shift int) string {

    shift = shift % 26

    if shift < 0 {
        shift += 26
    }


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
    // Test 1: Basic encryption/decryption
    text1 := "Hello, World!"
    encrypted1 := CaesarCipher(text1, 3)
    decrypted1 := CaesarCipher(encrypted1, -3)
    fmt.Printf("Test 1: %v\n", decrypted1 == text1)
    
    // Test 2: Large shift
    text2 := "abc"
    encrypted2 := CaesarCipher(text2, 30) // 30 % 26 = 4
    fmt.Printf("Test 2: 'abc' shifted by 30 = '%s'\n", encrypted2)
    
    // Test 3: What does shift 26 do?
    text3 := "hello"
    encrypted3 := CaesarCipher(text3, 26)
    fmt.Printf("Test 3: 'hello' shifted by 26 = '%s'\n", encrypted3)
    
    // Test 4: What does shift 0 do?
    text4 := "hello"
    encrypted4 := CaesarCipher(text4, 0)
    fmt.Printf("Test 4: 'hello' shifted by 0 = '%s'\n", encrypted4)
}

// func main() {
//     text := "Hello, World!"
//     shift := 3
    
//     encrypted := CaesarCipher(text, shift)
//     fmt.Printf("Original:  %s\n", text)
//     fmt.Printf("Encrypted: %s\n", encrypted)
    
//     // Try to decrypt
//     decrypted := CaesarCipher(encrypted, -shift)
//     fmt.Printf("Decrypted: %s\n", decrypted)

// }