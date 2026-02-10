# Exercise 1: The Palindrome Challenge

## Step 1: Independent Solution

### 1. Pseudocode
```
FUNCTION CheckPalindrome(input_string)
    IF length of input_string is 0 THEN
        RETURN false
    END IF
    FOR i from 0 to (length of input_string / 2) DO
        IF character_at_start is not equal to character_at_end THEN
            RETURN false 
        END IF
    END FOR
    RETURN true
END FUNCTION
```

### 2. Go Implementation
```go
package main
import "fmt"

// Checkpalindrome confirms whether a string reads the same forward and backward
func Checkpalindrome(arg string) bool {
    // empty strings are not considered palindromes
    if len(arg) == 0 {
        return false
    }
    // Compare characters from both ends moving toward the center
    // Only need to check half the string since we're comparing pairs
    for i := 0; i < len(arg)/2; i++ {
        // Compare character at position i from start 
        // with character at position i from end (len(arg)-1-i)
        if arg[i] != arg[len(arg)-1-i] {
            return false  // Characters don't match - not a palindrome
        }
    }
    // All character pairs matched - it's a palindrome
    return true
}

func main() {
    fmt.Println(Checkpalindrome("racecar"))   // Output: true
    fmt.Println(Checkpalindrome("hello"))     // Output: false
    fmt.Println(Checkpalindrome("A man a plan a canal Panama")) // Output: false
}
```

### 3. Test Results

**Test Case 1:** `"racecar"`
- Output: `true`
- Reason: "racecar" reads the same forwards and backwards

**Test Case 2:** `"hello"`
- Output: `false`
- Reason: "hello" backwards is "olleh" - not the same

**Test Case 3:** `"A man a plan a canal Panama"`
- Output: `false`
- Reason: Current implementation doesn't handle:
  - Case differences ('A' ≠ 'a')
  - Spaces between words
- These limitations will be addressed in Exercise 2

---

## Step 2: Using AI to Learn

### AI Analysis of My Palindrome Function

#### 1. Time Complexity
The function has **O(n)** time complexity, where n is the length of the string.

**Why?**
- The loop runs `len(arg)/2` times
- Even though it's half the string, in Big O notation we drop constants: n/2 = O(n)
- Each iteration does constant-time work: comparing two characters O(1)
- **Total:** O(n)

**Space Complexity:** O(1) — only uses a single variable `i`, with no extra space that grows with input size.

#### 2. Edge Cases I Might Miss

**✅ Already handled:**
- Empty string (returns false)

**Potential issues:**
- **Single character strings:** Input `"a"` returns `true` (correct - a single character is technically a palindrome)
- **Case sensitivity:** Input `"Racecar"` returns `false` because `'R' != 'r'` - most real-world checkers ignore case
- **Whitespace and punctuation:** Input `"A man a plan a canal Panama"` returns `false` because of spaces - typically checkers ignore non-alphanumeric characters
- **Unicode/special characters:** Input like `"café"` works correctly in Go because Go strings are UTF-8, but worth testing
- **Very long strings:** The algorithm handles this efficiently (O(n)) with no stack overflow risk since there's no recursion

#### 3. Better Approaches in Go

The two-pointer approach I used is already one of the best for this problem because it is:
- Efficient: O(n) time, O(1) space
- Clear and readable
- Optimal: You can't do better than O(n) because every character must be examined at least once

**Alternative approaches exist but are less efficient:**
- String reversal approach: O(n) time but O(n) space (creates new string)
- Recursive approach: O(n) time but O(n) space (call stack overhead)

My two-pointer approach is optimal for this problem.

---

## Step 3: Reflection

### What did I learn from solving it before asking AI?

I gained a practical understanding of two-pointer logic—specifically how to use a single loop to compare characters from both ends of a string simultaneously. By calculating the opposite index using `len(arg)-1-i`, I learned how to verify symmetry without needing a second loop or creating additional data structures.

### How is my understanding different now?

My understanding has shifted from simply "looping through text" to understanding efficiency and edge cases. I now see that while the core logic is simple, a robust program must account for factors like case-sensitivity (ASCII values) and memory management (comparing bytes in place vs. creating a new reversed string). I also learned about Big O notation and why n/2 operations still count as O(n) time complexity.

### Could I now write similar functions (e.g., reverse a string) without help?

Yes. Since I now understand how to access and compare mirror-image indices in a string, I can apply this same approach to reverse a string or rearrange elements in a collection. The pattern of using `i` and `len-1-i` is applicable to many similar problems, and I feel confident I could implement variations without relying on AI or external libraries.