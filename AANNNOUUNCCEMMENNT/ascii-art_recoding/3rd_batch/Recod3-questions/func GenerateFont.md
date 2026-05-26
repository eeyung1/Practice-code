Question 3: Programmatic Font Generator


Write a function GenerateFont that creates a complete mapping of printable ASCII characters (32-126) to their 8-line ASCII art representation using a mathematical/algorithmic approach instead of storing all characters. The generated font should use *, ., and space characters.

Requirements:

Function signature: func GenerateFont() map[rune][]string

Generate all 95 printable ASCII characters (space to tilde)

Each character should be 8 lines tall and 8 characters wide

Use algorithms based on character properties (vowels vs consonants, ascenders, descenders, etc.)

Characters should be somewhat distinguishable

Include proper spacing rules

File Structure:

ascii-art/
├── main.go
├── font/
│   ├── font.go
│   └── font_test.go


## test File

```go
package font

import (
    "testing"
)

func TestGenerateFont_Complete(t *testing.T) {
    font := GenerateFont()
    
    t.Run("Check character count", func(t *testing.T) {
        expectedCount := 95 // Space through tilde
        if len(font) != expectedCount {
            t.Errorf("Expected %d characters, got %d", expectedCount, len(font))
        }
    })
    
    t.Run("Check all printable ASCII", func(t *testing.T) {
        for c := rune(32); c <= rune(126); c++ {
            if _, exists := font[c]; !exists {
                t.Errorf("Missing character: %c (ASCII %d)", c, c)
            }
        }
    })
    
    t.Run("Check dimensions", func(t *testing.T) {
        for char, lines := range font {
            if len(lines) != 8 {
                t.Errorf("Character %c has %d lines, expected 8", char, len(lines))
            }
            for i, line := range lines {
                if len(line) != 8 {
                    t.Errorf("Character %c line %d has length %d, expected 8", char, i, len(line))
                }
            }
        }
    })
    
    t.Run("Verify character differences", func(t *testing.T) {
        // Different characters should have different representations
        a := font['A']
        b := font['B']
        
        identical := true
        for i := 0; i < 8; i++ {
            if a[i] != b[i] {
                identical = false
                break
            }
        }
        if identical {
            t.Error("Characters A and B have identical representations")
        }
    })
    
    t.Run("Space character should be all spaces", func(t *testing.T) {
        space := font[' ']
        for i, line := range space {
            for _, ch := range line {
                if ch != ' ' {
                    t.Errorf("Space char line %d contains non-space character: %c", i, ch)
                }
            }
        }
    })
}

func TestGenerateFont_VowelsSpecial(t *testing.T) {
    font := GenerateFont()
    
    // Vowels should be different from a sample consonant
    vowels := []rune{'A', 'E', 'I', 'O', 'U'}
    consonant := font['B']
    
    for _, vowel := range vowels {
        vowelRep := font[vowel]
        isDifferent := false
        
        for i := 0; i < 8; i++ {
            if vowelRep[i] != consonant[i] {
                isDifferent = true
                break
            }
        }
        
        if !isDifferent {
            t.Errorf("Vowel %c is identical to consonant B", vowel)
        }
    }
}




```