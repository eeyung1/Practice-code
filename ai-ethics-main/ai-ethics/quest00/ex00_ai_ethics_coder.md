## Exercise 00 — Learning to Use AI Ethically as a Coder
**Part A — Self-Assessment**
How I Have Used AI So Far

So far, I have mostly used AI to help me when I get stuck. Sometimes I ask it to explain errors and bugs in code . Other times, I directly ask it to write functions for me like how i want to creat a function for myself i use ai as when i dont know it and i want better explanation. If I’m being honest, there have been moments where I asked for the solution before fully trying on my own.

Yes, there were times I asked AI for solutions before struggling enough. Especially when I felt tired or frustrated as cs student dugging can be very stressful .

**Can I explain all the code I’ve submitted without AI?**

Not always. Sometimes I understand the general idea but not every line deeply which i think is a bad practics.

**If AI was suddenly unavailable during an exam** or interview, I would struggle in areas where I relied too much on it. I would feel less confident explaining complex logic or debugging under pressure.

## Which learner am I?

**Right now, I am somewhere between Learner A and Learner B.**
I want to say I’m Learner B, but honestly, I still sometimes act like Learner A when I rush.

## Where I want to be
I want AI to be a tool that sharpens my thinking and make faster than the 90s and make be fast learner , not something that replaces it. I want to struggle first, think deeply, and then use AI to improve my understanding—not avoid it.

## Part B — Palindrome Implementation
**Step 1: My Independent Attempt**
Pseudocode

Take the string.

Clean it (remove spaces, make lowercase if needed).

Compare the string to its reverse.

If they are equal → palindrome.

If not → not a palindrome.

My Python Implementation
def is_palindrome(text):
    cleaned = ""
    
    for char in text:
        if char.isalnum():
            cleaned += char.lower()
    
    reversed_text = cleaned[::-1]
    
    if cleaned == reversed_text:
        return True
    else:
        return False


# Testing
print(is_palindrome("racecar"))
print(is_palindrome("hello"))
print(is_palindrome("A man a plan a canal Panama"))

## Reflection

Struggling first helped me understand:

Why cleaning input matters.

Why case sensitivity affects results.

How slicing works in Python.

That reversing a string creates a new copy (memory cost).

If I had just copied the solution, I wouldn’t understand why isalnum() is important or why lowercase conversion is necessary.

Now I can confidently

Reverse a string

Check for symmetry

Modify logic to handle variations

# The mental model I built
A palindrome is simply a symmetry problem. If both ends match while moving inward, it’s valid.

**Part C — Variations**

To ignore punctuation and spaces, I used isalnum().

To make it case-insensitive, I used .lower().

If I wanted to return the position where it fails:

def palindrome_failure_position(text):
    cleaned = ""
    
    for char in text:
        if char.isalnum():
            cleaned += char.lower()
    
    left = 0
    right = len(cleaned) - 1
    
    while left < right:
        if cleaned[left] != cleaned[right]:
            return left
        left += 1
        right -= 1
    
    return -1


**This helped me understand two-pointer techniques**
 which are useful in many algorithm problems and remember an algorithm is a way of finding step by stey way of solving problem.

**Part D — My AI Fairness Contract**
I will use AI when

I have tried a problem seriously for at least 30 minutes.

I want to understand why my solution works or fails before trying to use ai to solve problem.

I want to compare approaches after building my own.

I want clarification on concepts.

I need help debugging after attempting it myself.

# I will NOT use AI when

I haven’t tried the problem

I am writing an exam or doing an assessment.

I need to strengthen foundational thinking.

I feel lazy and just want a shortcut.

I cannot explain the problem in my own words.

# I know I’m using AI fairly when

I can explain my code clearly without looking.

I can solve similar problems alone.

I feel more confident, not dependent.

I understand the (why,” not just the “how.)

**Signed: Ejike Victor**
**Date: 14 February 2026**

# Part E — Real-World Scenario

If I rely too much on AI, I won’t survive real-world pressure thre problem in  a situation why system brackdown.

In an interview, if someone asks me to design a caching system, I must understand:

**What caching is**

Why it improves performance

Memory trade-offs

Expiration strategies

If production breaks at 2 AM and AI is unavailable, I must be able to debug code logically.

If a new technology has little documentation, I must rely on experimentation, reading docs, and reasoning.

Using AI fairly now trains me to think independently. It builds confidence, problem-solving ability, and resilience. AI should support my thinking—not replace it.

# Part F — Skills Assessment
Self-Rating (1–5)

Problem decomposition: 3
Systems thinking: 5
Critical evaluation: 3
Debugging mindset: 4
Conceptual understanding (the “why”): 3

Lowest area Systems thinking 

Improvement Plan

**This week I will**

Break down at least 3 coding problems into diagrams before coding.

Read documentation instead of asking AI immediately.

Debug one issue fully on my own before searching for help.

Goal: Improve systems thinking by understanding how parts connect, not just how functions work.