# Exercise 2: Simulated API Call (Weather Example)

## Objective
Practice designing prompts that mimic function calls by defining JSON schemas and testing outputs across multiple inputs.

---

## Step 1: Unconstrained Simulated API Prompt

**Prompt:**
> "Simulate a weather API response for Paris today."

**Model Output:**

```json
{
  "city": "Paris",
  "country": "FR",
  "date": "2026-02-22",
  "weather": {
    "description": "Partly cloudy",
    "temperature_celsius": 11,
    "temperature_fahrenheit": 51.8,
    "feels_like_celsius": 10,
    "humidity_percent": 68,
    "wind": {
      "speed_kmh": 15,
      "direction": "NW"
    },
    "pressure_hPa": 1015,
    "visibility_km": 10
  },
  "forecast": {
    "morning": {
      "temperature_celsius": 9,
      "description": "Cloudy"
    },
    "afternoon": {
      "temperature_celsius": 12,
      "description": "Partly cloudy"
    },
    "evening": {
      "temperature_celsius": 10,
      "description": "Clear"
    }
  }
}
```

> **Observation:** Without a defined schema, the model returned nested objects, forecast segments, and multiple fields it decided on its own — far more than may be needed by a real application.

---

## Step 2: Schema-Constrained Simulated API Prompt

**Defined Schema:**
```json
{
  "location": string,
  "temperature_celsius": number,
  "condition": string
}
```

**Prompt:**
> "Simulate a weather API response for Paris today. Return only the following JSON schema: {"location": string, "temperature_celsius": number, "condition": string}."

**Model Output:**

```json
{
  "location": "Paris",
  "temperature_celsius": 11,
  "condition": "Partly cloudy"
}
```

✅ Output matches the defined schema exactly — no extra fields, correct data types.

---

## Step 3: Testing with Multiple Cities

### Test 2 — London

**Prompt:**
> "Simulate a weather API response for London today. Return only the following JSON schema: {"location": string, "temperature_celsius": number, "condition": string}."

**Model Output:**

```json
{
  "location": "London",
  "temperature_celsius": 9,
  "condition": "Light rain"
}
```

✅ Schema held — same three fields, correct data types.

---

### Test 3 — New York

**Prompt:**
> "Simulate a weather API response for New York today. Return only the following JSON schema: {"location": string, "temperature_celsius": number, "condition": string}."

**Model Output:**

```json
{
  "location": "New York",
  "temperature_celsius": 4,
  "condition": "Sunny"
}
```

✅ Schema held — consistent structure across all three cities.

---

## Step 4: Comparison Analysis

The free output generated information that had more details and so much unnecessary information while the schema-defined output was precise and generated the required information. A developer would prefer the schema constrained version because it provides definite outputs and eliminates unnecessary information. Testing with multiple cities generated the expected output which was precise and exact.

---

## Summary Table

| Aspect | Unconstrained Prompt | Schema-Constrained Prompt |
|---|---|---|
| Fields Returned | Many — nested objects, forecasts | Exactly 3 — location, temperature, condition |
| Schema Control | None — model decided | Enforced — defined by prompt |
| Output Predictability | Low | High |
| Consistent Across Cities | N/A | Yes — verified across Paris, London, New York |
| Suitable for Real API Use | No | Yes |
