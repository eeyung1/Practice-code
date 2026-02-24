# Exercise 3: Chaining Prompts for Agent Flow

## Objective
Combine multiple prompts in sequence to simulate a tool-using agent, where each prompt feeds into the next to mimic a real AI workflow.

---

## Overview: The Agent Assembly Line

| Step | Action | Input | Output |
|---|---|---|---|
| Step 1 | Capture user query | Natural language question | Structured JSON intent |
| Step 2 | Simulate API call | JSON intent | Weather data JSON |
| Step 3 | Format final answer | Weather data JSON | Human-friendly response |

---

## Step 1: Capture the User Query

**Prompt:**
> "A user asks: What is the temperature in Lagos today? Extract the intent and location as JSON: {"intent": string, "location": string}."

**Model Output:**

```json
{
  "intent": "get_temperature",
  "location": "Lagos"
}
```

> **Observation:** The model correctly extracted the intent and location from a natural language question and structured it as machine-readable JSON, giving the agent a clear understanding of what the user wants and where.

---

## Step 2: Simulate the API Call

**Prompt:**
> "Given the following input: {"intent": "get_temperature", "location": "Lagos"}, simulate a weather API response. Return only: {"location": string, "temperature_celsius": number, "condition": string}."

**Model Output:**

```json
{
  "location": "Lagos",
  "temperature_celsius": 32,
  "condition": "Partly cloudy"
}
```

> **Observation:** The agent used the structured output from Step 1 as input, simulated an API call, and returned clean, schema-constrained weather data ready for the next step.

---

## Step 3: Format the Final Answer for the User

**Prompt:**
> "Given this weather data: {"location": "Lagos", "temperature_celsius": 32, "condition": "Partly cloudy"}, write a friendly one-sentence response to the user who asked: What is the temperature in Lagos today?"

**Model Output:**

It's currently about 32°C in Lagos, with partly cloudy skies — a warm day, so stay cool!

> **Observation:** The structured JSON data from Step 2 was converted back into a natural, human-friendly response, completing the full agent loop from question to answer.

---

## Step 4: Comparison Analysis

The first step captured the user query, the second simulated the weather API call using the location extracted, and the third turned it back into a natural, human-friendly response. Breaking it into steps ensures clarity and reasoning flow, with every step spelt out for efficiency. When real AI agents are required to perform a certain task, they follow a sequence of steps demonstrated by the work done above, which is the same logic used by applications like Siri, Google Assistant, and customer support chatbots to understand a request, fetch data, and deliver a response.

---

## Full Chain Summary

```
User Question (natural language)
        ↓
[Step 1] Extract intent + location → JSON
        ↓
[Step 2] Simulate API call → Weather JSON
        ↓
[Step 3] Format response → Human-friendly answer
        ↓
Final Output delivered to user
```

---

## Summary Table

| Aspect | Single Prompt | Chained Prompts |
|---|---|---|
| Clarity | Low — all mixed together | High — each step is isolated |
| Reasoning Flow | Uncontrolled | Structured and sequential |
| Output Control | Unpredictable | Precise at each stage |
| Real-world Applicability | Limited | Mirrors how real AI agents work |
| Debugging Ease | Difficult | Easy — each step can be tested independently |
