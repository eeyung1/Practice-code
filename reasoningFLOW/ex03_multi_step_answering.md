
# Exercise 3: Multi-Step Question Answering

## Objective
Break complex questions into structured steps and compare basic prompting with structured prompting.

---

## Selected Problem

Write a Python function that takes 10 numbers and returns the sum of the first five and the sum of the last five multiplied together.

---

## Prompt 1 (Basic Prompt)

Write a Python function that takes 10 numbers and returns the sum of the first five and the sum of the last five multiplied together.

### Model Output

```python
def calculate_product(numbers):
    if len(numbers) != 10:
        raise ValueError("The list must contain exactly 10 numbers.")

    first_five_sum = sum(numbers[:5])
    last_five_sum = sum(numbers[5:])

    return first_five_sum * last_five_sum
```

### Brief Analysis

The first prompt gives a direct solution without showing reasoning. It focuses only on the final implementation.

---

## Prompt 2 (Structured Multi-Step Prompt)

Write a Python function that takes 10 numbers and returns the sum of the first five and the sum of the last five multiplied together.  

Before giving the final answer:  
1. Restate the problem.  
2. Break it into smaller logical steps.  
3. Solve each step.  
4. Then provide the final Python function.

### Model Output

Restating the problem: The task is to write a function that accepts 10 numbers, computes the sum of the first five and the sum of the last five, then multiplies both results.

Step 1: Validate that the input contains exactly 10 numbers.  
Step 2: Compute the sum of the first five numbers.  
Step 3: Compute the sum of the last five numbers.  
Step 4: Multiply both sums and return the result.

```python
def calculate_product(numbers):
    if len(numbers) != 10:
        raise ValueError("The list must contain exactly 10 numbers.")

    first_five_sum = sum(numbers[:5])
    last_five_sum = sum(numbers[5:])

    result = first_five_sum * last_five_sum
    return result
```

### Brief Analysis

The structured prompt improves clarity by forcing the model to explain its reasoning before giving the final code. It improves transparency and makes debugging easier because each logical step is clearly defined.

---

## Comparison Summary

The basic prompt returns a direct solution, while the structured prompt encourages explicit reasoning and decomposition. This demonstrates how controlled prompting improves clarity and problem-solving structure in AI-assisted development.
