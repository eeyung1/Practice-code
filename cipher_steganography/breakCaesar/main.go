package main

import (
	"fmt"
)

func breakCaesar(ciphertext string) string {
    freq := countFrequencies(ciphertext)
    mostCommon := mostCommonLetter(freq)
    shift := findShift(mostCommon)
    return CaesarCipher(ciphertext, -shift)  
}

func main() {
    original := `the quick brown fox jumps over the lazy dog the quick brown fox jumps over the lazy dog the quick brown fox jumps over the lazy dog`
    shift := 3
    encrypted := CaesarCipher(original, shift)

    fmt.Println(bruteForceCaesar(encrypted))
    cracked := breakCaesar(encrypted)
    _ = cracked
    
    // fmt.Printf("Original:  %s\n", original[:50])
    // fmt.Printf("Cracked:   %s\n", cracked[:50])
    // fmt.Printf("Success:   %v\n", cracked == original)
}