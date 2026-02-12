# Exercise 1: The Palindrome Challenge

## Step 1: Independent Solution

### 1. Pseudocode
```
FUNCTION CheckPalindrome(input_string)
    IF length of input_string is 0 THEN
        RETURN true
    END IF
    FOR i from 0 to (length of input_string / 2) DO
        IF character_at_start is not equal to character_at_end THEN
            RETURN false 
        END IF
    END FOR
    RETURN true
END FUNCTION
```

### 2.python Implementation
```python
def check_palindrome(s):
    if len(s) == 0:
        return True #An empty string is a palindrome so this line of logic tells the computer to do just that.  

    for i in range(len(s) // 2): #The computer has to go through each character in the string so as to compare
        if s[i] != s[len(s)-1-i]: #While going through the characters, it compares the first and the last characte, the second and second to the last and so on.
            return False  #if it doesnot match, it returns false (not a palindrone)

        return True  #if it matches, it returns true(is a palindrome)
    
print(check_palindrome("raecar"))
print(check_palindrome("hello"))
print(check_palindrome("A man a plan a canal Panama"))

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

# Palindrome Function Analysis

This document evaluates a Python implementation of a palindrome checker, covering its complexity, potential pitfalls, and alternative methods.

---

### 1. Time and Space Complexity
For an input string of length $n$:

*   **Time Complexity:** $O(n)$
    Even though the loop only iterates up to $n/2$, Big O notation ignores constant factors. The function performs a linear scan, making it highly efficient.
*   **Space Complexity:** $O(1)$
    The function is memory-efficient because it does not create additional data structures (like lists or copies of the string) that scale with the input size. It only uses a few integer variables for indexing.

---

### 2. Edge Cases to Consider
While the current logic handles basic lowercase strings, it may fail in real-world scenarios:

*   **Case Sensitivity:** "Racecar" will return `False` because the ASCII values of `'R'` and `'r'` are different.
*   **Non-Alphanumeric Characters:** Phrases like "A man, a plan, a canal, Panama" will fail due to spaces and commas.
*   **Empty or Single-Character Strings:** Your code correctly returns `True` for `len == 0`. Single-character strings will also pass (as the loop range becomes zero), which is the correct behavior.

---

### 3. Better Approaches & Refinement

#### The "Pythonic" Way (Slice)
If memory is not a strict constraint, the slicing method is the most concise. It is highly optimized in Python because the reversal happens at the C-level.

```python
def is_palindrome(s):
    # Returns True if the string is identical to its reverse
    return s == s[::-1]
```

## Step 3: Reflection

### What did I learn from solving it before asking AI?

I gained a practical understanding of two-pointer logic—specifically how to use a single loop to compare characters from both ends of a string simultaneously. 

### How is my understanding different now?

My understanding has shifted from simply "looping through text" to understanding efficiency and edge cases. I now see that while the core logic is simple, a robust program must account for factors like case-sensitivity (ASCII values) and memory management (comparing bytes in place vs. creating a new reversed string). I also learned about Big O notation and why n/2 operations still count as O(n) time complexity.

### Could I now write similar functions (e.g., reverse a string) without help?

Yes. Since I now understand how to access and compare mirror-image indices in a string, I can apply this same approach to reverse a string or rearrange elements in a collection. 