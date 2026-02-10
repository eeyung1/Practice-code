# Exercise 2: Strengthening Through Variations

## Step 1: Modifying the Palindrome Function

The palindrome function was extended to handle more cases by:
- Ignoring spaces and punctuation  
- Making the comparison case-insensitive  
- Returning the position where the string stops being a palindrome, or `-1` if it is a palindrome  

### Implementation (Go)

```go
package main

import (
    "fmt"
    "unicode"
    "strings"
)

// cleanString removes spaces, punctuation, and converts characters to lowercase
func cleanString(s string) string {
    cleaned := ""
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            cleaned += string(r)
        }
    }
    return strings.ToLower(cleaned)
}

// Checkpalindrome returns -1 if the string is a palindrome,
// or the position where the palindrome check fails
func Checkpalindrome(arg string) int {
    if len(arg) == 0 {
        return 0
    }

    cleaned := cleanString(arg)

    for i := 0; i < len(cleaned)/2; i++ {
        if cleaned[i] != cleaned[len(cleaned)-1-i] {
            return i
        }
    }

    return -1
}

func main() {
    fmt.Println(Checkpalindrome("racecar"))                        // -1
    fmt.Println(Checkpalindrome("A man, a plan, a canal: Panama")) // -1
    fmt.Println(Checkpalindrome("hello"))                          // 0
}
```

---

## Step 2: AI Evaluation

### Did I miss anything?

- The palindrome comparison uses byte indexing instead of rune indexing, which can cause failures for non-ASCII characters (e.g., accented letters).
- The returned break position refers to the cleaned string, not the original input. This behavior is valid but undocumented.
- Empty or non-alphanumeric strings return `-1` (palindrome). This is reasonable but implicit.

### Can it be more efficient?

- Using `+=` for string concatenation inside a loop can be inefficient and lead to O(n²) behavior.
- A two-pointer approach that skips non-alphanumeric characters during comparison would avoid creating a cleaned string and reduce memory usage.

---

## Step 3: Reflection — What AI Added That I Didn’t Consider Initially

AI highlighted that the solution could fail for non-ASCII characters due to byte-based string indexing. It also pointed out that while returning `-1` for a palindrome is reasonable, this behavior is undocumented and could be clarified. Additionally, AI showed how the solution could be made more efficient by avoiding repeated string concatenation and unnecessary memory allocation.