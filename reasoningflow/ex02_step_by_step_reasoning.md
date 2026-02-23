# Exercise 2: Step-by-Step Reasoning Prompt

## Objective

Use structured prompting to encourage reasoning before giving a final
answer and compare outputs with and without chain-of-thought
instructions.

------------------------------------------------------------------------

## Selected Problem

Write a Python function that checks if a number is prime.

------------------------------------------------------------------------

## Prompt 1 (Without Reasoning Instruction)

Write a Python function that checks if a number is prime.

### Model Output

``` python
def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, n):
        if n % i == 0:
            return False
    return True
```

### Brief Analysis

The model provides only the final code without explaining the reasoning
process. The logic used to determine whether the number is prime is not
explicitly described.

------------------------------------------------------------------------

## Prompt 2 (With Step-by-Step Reasoning Instruction)

Write a Python function that checks if a number is prime. Let's think
step by step before answering.

### Model Output

Step 1: A prime number is a number greater than 1 that has no divisors
other than 1 and itself.

Step 2: If the number is less than or equal to 1, it is not prime.

Step 3: Check divisibility from 2 up to n-1. If the number is divisible
by any of these numbers, it is not prime.

Step 4: If no divisors are found, then the number is prime.

``` python
def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, n):
        if n % i == 0:
            return False
    return True
```

### Brief Analysis

The main structural difference is that the second prompt makes the
reasoning process visible before presenting the final code. The
step-by-step approach improves transparency and reduces the risk of
logical mistakes because each step can be verified. This method is
especially useful for debugging and learning since errors can be
identified more easily by reviewing each intermediate step.

------------------------------------------------------------------------

## Comparison Summary

The version without reasoning produces a direct solution, while the
step-by-step version explains the logic before coding. Chain-of-thought
prompting increases clarity, transparency, and reliability of the
model's output.
