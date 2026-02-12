# Exercise 2: Strengthening Through Variations

## Step 1: Modifying the Palindrome Function

The palindrome function was extended to handle more cases by:
- Ignoring spaces and punctuation  
- Making the comparison case-insensitive  
- Returning the position where the string stops being a palindrome, or `-1` if it is a palindrome  

### Implementation (python)

```python
def check_palindrome(s):
    clean = ''.join(c.lower() for c in s if c.isalnum())
    
    for i in range(len(clean) // 2):  # Compare characters from both ends moving toward center
        if clean[i] != clean[-(i+1)]: # checking if the both characters are either the same or different
            return i  # Position where palindrome breaks
    
    return -1  # String is a palindrome
```

---

### Test Results

| Input | Output | Reason |
|-------|--------|--------|
| `"racecar"` | `-1` | Is a palindrome |
| `"hello"` | `0` | Breaks at position 0: `'h' != 'o'` |
| `"A man a plan a canal Panama"` | `-1` | Is a palindrome after cleaning spaces |

---

## Step 2: AI Feedback

**Question asked:**
> "I modified my palindrome function to handle more cases. Did I miss anything? Can it be more efficient?"

**AI Response Summary:**

- The two-pointer approach using negative indexing `clean[-(i+1)]` is clean and efficient
- Time complexity is **O(n)** — loop runs `n/2` times, constants are dropped in Big O notation
- Space complexity is **O(n)** — `clean` creates a new string in memory
- Possible edge cases to consider:
  - Strings with only spaces/punctuation (e.g., `"!!!"`) → `clean` becomes `""` → returns `-1` (handled correctly since loop never runs)
  - Unicode/accented characters (e.g., `"café"`) → `isalnum()` handles these correctly in Python
  - Numbers in strings (e.g., `"12321"`) → treated as alphanumeric, works correctly
---

## Step 3: Reflection

**What did AI add that I didn't consider initially?**

The edge case of a string containing only spaces or punctuation like `"!!!"` was something I hadn't thought about. After cleaning, it becomes an empty string, and since `range(0)` produces no iterations, the function returns `-1` — treating it as a palindrome. Whether that's the right behavior depends on the use case. AI also confirmed that my use of negative indexing `-(i+1)` is equivalent to `len(clean)-1-i` but more Pythonic, which I verified by tracing through the logic manually before using it.