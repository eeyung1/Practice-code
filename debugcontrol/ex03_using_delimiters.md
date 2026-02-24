# Exercise 3: Using Delimiters to Define Context

## Objective
Restrict responses to a specific block of text using delimiters, and compare outputs with and without them.

---

## Step 1: Prompt WITHOUT Delimiters

**Prompt:**
```
Here is some text: The Amazon rainforest produces 20% of the world's oxygen...

Question: What does the text say about the Amazon?
```

**Model Output:**

The text says that the Amazon rainforest produces 20% of the world's oxygen.

---

## Step 2A: Prompt WITH Incorrectly Placed Delimiters

**Prompt:**
```
###
The Amazon rainforest is often called the "lungs of the Earth" and credited with producing some 20 percent of the world's oxygen.###
But that's often a repeated, yet misleading claim. In fact, it's a gross overestimate. The Amazon's net contribution to the oxygen we breathe likely hovers around zero.

Answer using only the text inside the delimiters: What does the text say about the Amazon?
```

**Model Output:**

The Amazon rainforest is often called the "lungs of the Earth" and credited with producing some 20 percent of the world's oxygen.

> **Note:** Because the closing `###` was placed in the middle of the passage, the model treated only the first sentence as the delimited content, ignoring the rest of the paragraph entirely.

---

## Step 2B: Prompt WITH Correctly Placed Delimiters

**Prompt:**
```
###
The Amazon rainforest is often called the "lungs of the Earth" and credited with 
producing some 20 percent of the world's oxygen. But that's often a repeated, yet 
misleading claim. In fact, it's a gross overestimate. The Amazon's net contribution 
to the oxygen we breathe likely hovers around zero.
###

Answer using only the text inside the delimiters: What does the text say about the Amazon?
```

**Model Output:**

The Amazon rainforest is often called the "lungs of the Earth" and credited with producing some 20 percent of the world's oxygen, but that's often a repeated, yet misleading claim; in fact, it's a gross overestimate, and the Amazon's net contribution to the oxygen we breathe likely hovers around zero.

---

## Step 3: Comparison Analysis

Without delimiters, the answer was direct and concise, but the model was free to draw from anywhere. When delimiters were incorrectly placed in the middle of the passage, the model only used the first sentence, which shows that delimiters directly specify the scope of text the AI is supposed to use to generate its output. When the delimiters were correctly placed wrapping the full passage, the model stayed entirely within that boundary and captured all the key points. In a real-world scenario such as a customer support bot, correct delimiter placement would ensure that answers are generated only within the scope of the company's FAQ document, preventing the model from going off topic. In cases where there is a bigger text with no delimiters, the scope of the output would outpace the text provided, making responses harder to control and verify.

---

## Summary Table

| Aspect | No Delimiters | Incorrect Delimiters | Correct Delimiters |
|---|---|---|---|
| Boundary Control | None | Partial — cut off mid-passage | Full — entire passage wrapped |
| Model Scope | Free to draw from anywhere | Only used first sentence | Stayed within full passage |
| Output Accuracy | Correct but uncontrolled | Incomplete | Complete and constrained |
| Real-world Reliability | Low | Low | High |
| Recommended? | No | No | Yes |
