# Exercise 1: Logic Puzzle with Explanation

---

## Objective
Guide the model to solve a problem with explicit reasoning by using chain-of-thought prompting.

---

## Logic Puzzle Used

"A mother is 30 years old, which is 9 years more than three times her daughter's age. Find the daughter's age."

---

## Test 1: WITH Reasoning Instruction

### Prompt

> "A mother is 30 years old, which is 9 years more than three times her daughter's age. Find the daughter's age.
>
> Explain your reasoning step by step before giving the final answer."

### Model Output

Let the daughter's age be x years.

**Step 1: Translate the statement into an equation**

The mother is 30 years old. This is 9 years more than three times her daughter's age.

So we write:
```
30 = 3x + 9
```

**Step 2: Solve the equation**

Subtract 9 from both sides:
```
30 - 9 = 3x
21 = 3x
```

Now divide both sides by 3:
```
x = 7
```

**Final Answer:** The daughter is 7 years old.

---

## Test 2: WITHOUT Reasoning Instruction

### Prompt

> "A mother is 30 years old, which is 9 years more than three times her daughter's age. Find the daughter's age."

### Model Output

The daughter is 7 years old. The mother was 30, and her daughter was 7, so there's a 23-year difference between them. You could even think about how the difference in their ages will always stay the same, even as they both get older!

---

## Comparison and Analysis

The first prompt clearly explains the reasoning behind each of the steps in getting to the answer. It sets up a variable, translates the problem into an algebraic equation, and shows the mathematical operations needed to isolate the variable and solve for x.

The first prompt is better for learning because it shows a step-by-step approach to how the solution is obtained instead of just giving the answer to the question. The second output provides the correct answer but offers no insight into the problem-solving process, making it useful only for verification rather than education.

**Key Insight:** Adding the instruction "Explain your reasoning step by step before giving the final answer" transforms the model from an answer generator into a teaching tool. The intermediate steps allow verification of logic, identification of potential errors, and understanding of the methodology rather than just memorizing the result.
