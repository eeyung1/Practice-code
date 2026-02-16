# Exercise 00: Learning to Use AI Ethically as a Coder

---

## Part A: Self-Assessment

**How have you used AI for coding so far?**
I have used AI in trying to understand the logic behind any function and then getting the code to see how it's executed.

**Do you ask AI for solutions before trying yourself?**
I use AI to guide me through the learning process of getting the solutions myself.

**Can you explain code you've submitted without AI's help?**
Definitely. I ensure I understand the code and the logic behind it before submitting anything.

**What would happen if AI was suddenly unavailable during an exam or interview?**
I am certain that with my love for understanding fundamentals, I will still do quite well.

**Which learner are you?**
I consider myself a Learner B.

**Where I am today and where I want to be:**
Currently I am in the habit of endlessly searching for the right and most efficient way of understanding the fundamentals of whatever I am doing. Where I want to be is to master the fundamentals of programming and get the computer to do whatever I want it to do.

---

## Part B: Palindrome Solution + Reflection

### Pseudocode
```
FUNCTION CheckPalindrome(input_string)
    IF length of input_string is 0 THEN
        RETURN True
    END IF
    FOR i from 0 to (length of input_string / 2) DO
        IF character_at_start is not equal to character_at_end THEN
            RETURN false
        END IF
    END FOR
    RETURN true
END FUNCTION
```

### Python Implementation
```python
def check_palindrome(s):
    if len(s) == 0:
        return False

    for i in range(len(s) // 2):
        if s[i] != s[-(i+1)]:
            return False

    return True


# Test cases
print(check_palindrome("racecar"))   # Output: True
print(check_palindrome("hello"))     # Output: False
print(check_palindrome("A man a plan a canal Panama"))  # Output: False
```

### Test Results

| Input | Output | Reason |
|-------|--------|--------|
| `"racecar"` | `True` | Reads the same forwards and backwards |
| `"hello"` | `False` | "hello" backwards is "olleh" |
| `"A man a plan a canal Panama"` | `False` | Spaces and case not handled yet |

### Step 2: Strategic AI Use

**Question asked to AI:**
> "I wrote a palindrome function in Python. Here's my code: [code above]. What's the time complexity? What edge cases might I miss? Are there better approaches?"

**AI Analysis Summary:**
- Time complexity is **O(n)** — loop runs n/2 times, constants dropped in Big O notation
- Space complexity is **O(1)** — no extra data structures created
- Edge cases identified: case sensitivity, spaces/punctuation, unicode characters, single characters
- Two-pointer approach (my approach) is optimal — better than string reversal which uses O(n) space

### Step 3: Reflection

**What did I learn from solving it before asking AI?**
I gained a practical understanding of two-pointer logic — specifically how to use a single loop to compare characters from both ends of a string simultaneously. By calculating the opposite index using negative indexing `-(i+1)`, I learned how to verify symmetry without needing a second loop or creating additional data structures.

**How is my understanding different now?**
My understanding has shifted from simply "looping through text" to understanding efficiency and edge cases. I now see that while the core logic is simple, a robust program must account for factors like case-sensitivity (ASCII values) and memory management — comparing bytes in place versus creating a new reversed string. I also learned about Big O notation and why n/2 operations still count as O(n) time complexity.

**Could I now write similar functions without help?**
Yes. Since I now understand how to access and compare mirror-image indices in a string using both `len(s)-1-i` and `-(i+1)`, I can apply this same approach to reverse a string or rearrange elements in a collection across different languages.

---

## Part C: Palindrome Variations

### Modified Python Implementation
```python
def check_palindrome(s):
    # Remove non-alphanumeric characters and convert to lowercase
    clean = ''.join(c.lower() for c in s if c.isalnum())

    # Compare characters from both ends moving toward center
    # Returns position where palindrome breaks, or -1 if palindrome
    for i in range(len(clean) // 2):
        if clean[i] != clean[-(i+1)]:
            return i  # Position where palindrome breaks

    return -1  # String is a palindrome


# Test cases
print(check_palindrome("racecar"))                     # Expected: -1 (palindrome)
print(check_palindrome("hello"))                       # Expected: 0 (breaks at position 0)
print(check_palindrome("A man a plan a canal Panama")) # Expected: -1 (palindrome after cleaning)
```

### Changes Made:
1. Added string cleaning using `isalnum()` — removes spaces and punctuation
2. Added `c.lower()` — makes comparison case-insensitive
3. Changed return type from `bool` to `int`:
   - Returns `-1` if string IS a palindrome
   - Returns position (0+) where it breaks if NOT a palindrome

### Test Results

| Input | Output | Reason |
|-------|--------|--------|
| `"racecar"` | `-1` | Is a palindrome |
| `"hello"` | `0` | Breaks at position 0: 'h' != 'o' |
| `"A man a plan a canal Panama"` | `-1` | Is a palindrome after cleaning spaces |

---

## Part D: Personal Fairness Contract

### I will use AI when:
- I have tried solving a problem myself for at least 20 minutes
- I need to understand WHY something works, not just HOW
- I want to explore alternative approaches after having a working solution
- I need to verify my understanding of a concept I've already attempted

### I will NOT use AI when:
- I haven't attempted the problem on my own first
- I am completing an exam, test, or assessment
- I am still building fundamental concepts from scratch
- I want a quick answer just to move on without understanding

### I know I'm using AI fairly when:
- I can explain my code completely without looking at AI's response
- I could solve a similar problem independently tomorrow
- I feel more confident in my own understanding after the interaction
- AI deepened my thinking rather than replaced it

**Signed:** Emmanuel Eyung
**Date:** 16th February 2026

---

## Part E: Real-World Scenario Analysis

**Interview: "Explain how you'd implement a caching system."**
If I had always relied on AI to design systems, I would struggle to explain my reasoning live without a keyboard in front of me. An interviewer isn't just looking for the right answer — they're looking for how you think. Over-reliance on AI means you've seen solutions but never built the mental model behind them. You'd freeze because there's nothing to draw from except memory of AI outputs you never truly understood.

**Production Bug at 2 AM: AI is unavailable.**
This is the scenario that exposes over-reliance most brutally. If AI wrote the code and I never understood it, debugging it alone at 2 AM becomes nearly impossible. I'd be reading code like a foreign language. Using AI fairly means I always understand what I submit — so when it breaks, I have a mental map of where to look and why something might be failing.

**New Technology with Little Documentation:**
If I never learned to read official documentation and experiment independently, a new library with no AI training data leaves me completely stuck. Using AI fairly builds the habit of reading docs, forming hypotheses, and testing them — skills that work regardless of whether AI knows the technology or not.

**How using AI fairly prepares me for these scenarios:**
Using AI as a learning amplifier rather than an answer generator means every interaction builds a mental model I actually own. When AI is unavailable — in interviews, emergencies, or with new technology — I'm not lost because I never outsourced my thinking. The struggle of attempting problems first is what builds the judgment and confidence that no AI outage can take away.

---

## Part F: Skills Assessment + Action Plan

| Skill | Description | Rating | Improvement Plan |
|-------|-------------|--------|-----------------|
| Problem Decomposition | Breaking down problems logically | 4/5 | Practice breaking every new problem into smallest possible steps before writing any code |
| Systems Thinking | Understanding how components interact | 3/5 | Study how real applications connect frontend, backend, database, and cloud together |
| Critical Evaluation | Knowing when code is wrong or inefficient | 4/5 | Always ask "is there a better way?" before submitting any solution |
| Debugging Mindset | Investigating unexpected behavior | 3/5 | Practice reading error messages carefully and forming hypotheses before changing code |
| Conceptual Understanding | Knowing WHY, not just HOW | 4/5 | For every concept learned, write a TIL entry explaining it in my own words |

**Lowest Rated Skills:** Systems Thinking and Debugging Mindset (both 3/5)

**3 Specific Actions This Week to Improve (Without AI):**

1. **Systems Thinking:** Draw a diagram of how my palindrome function would work as an API — what happens from the moment a user sends a request to when they receive a response. Do this without AI, just pen and paper.

2. **Debugging Mindset:** When my code breaks this week, write down my hypothesis of what's wrong BEFORE looking at the error message or changing anything. Then compare my prediction with the actual error.

3. **Both Skills Combined:** Take one real application I use daily (GitHub, WhatsApp, etc.) and write a one-paragraph explanation of how I think it works internally — what components exist, how they talk to each other, where bugs might hide.
