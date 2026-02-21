# Exercise 2: Role-Based Mentoring and Feedback

## Buggy Code Submitted

``` python
def reverse(input)
reverse = [:-1]
return reverse
```

## Prompt Used

Act as a coding mentor and review my code by providing corrections and
explain improvements clearly.

## Model Output

Thank you for sharing your function. I can see you're trying to reverse
a string using slicing, which is a good approach. However, there are a
few issues that need correction.

First, the function definition is missing a colon at the end. Second,
the slicing syntax is incorrect --- to reverse a string in Python, you
need to use `input[::-1]`. Also, the variable name `reverse` should not
be the same as the function name, as this can cause confusion. Finally,
proper indentation is required inside the function.

Here is a corrected version:

``` python
def reverse(text):
    reversed_text = text[::-1]
    return reversed_text
```

This version clearly separates the function name and variable name, uses
correct slicing syntax, and follows Python's indentation rules. Keep
practicing --- you're on the right track.

## Analysis

The mentor tone is supportive by encouraging my effort before explaining
the mistakes and making the necessary corrections. The response does not
immediately criticize the errors but acknowledges the intention behind
the code.

Supportive language such as "I can see you're trying to reverse a string
using slicing, which is a good approach" shows encouragement. The role
instruction influenced the feedback style because the explanation was
constructive and patient rather than blunt or overly technical.
