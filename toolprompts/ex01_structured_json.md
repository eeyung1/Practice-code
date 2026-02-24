# Exercise 1: Structured JSON Output

## Objective
Create prompts that produce structured machine-readable output and verify consistency across multiple inputs.

---

## Step 1: Unstructured Prompt

**Prompt:**
> "List 5 continents with their number of countries."

**Model Output:**

Here is a list of 5 continents along with the approximate number of recognized countries in each:

1. Africa – 54 countries
2. Asia – 49 countries
3. Europe – 44 countries
4. North America – 23 countries
5. South America – 12 countries

---

## Step 2: Structured JSON Prompt

**Prompt:**
> "List 5 continents with their number of countries. Return the result in JSON format: [{"continent": string, "number_of_countries": number}]."

**Model Output:**

```json
[
  {"continent": "Africa", "number_of_countries": 54},
  {"continent": "Asia", "number_of_countries": 49},
  {"continent": "Europe", "number_of_countries": 44},
  {"continent": "North America", "number_of_countries": 23},
  {"continent": "South America", "number_of_countries": 12}
]
```

✅ JSON is valid — proper brackets, quoted keys, correct data types, no missing commas.

---

## Step 3: Testing with Multiple Inputs

### Test 2 — Planets and Distance from the Sun

**Prompt:**
> "Return the 5 planets closest to the sun with their distance from the sun in JSON format: [{"planet": string, "distance_from_sun_km": number}]."

**Model Output:**

```json
[
  {"planet": "Mercury", "distance_from_sun_km": 57900000},
  {"planet": "Venus", "distance_from_sun_km": 108200000},
  {"planet": "Earth", "distance_from_sun_km": 149600000},
  {"planet": "Mars", "distance_from_sun_km": 227900000},
  {"planet": "Jupiter", "distance_from_sun_km": 778500000}
]
```

✅ JSON is valid — different field names, large numeric values, structure held consistently.

---

### Test 3 — Nigerian States and Estimated Population

**Prompt:**
> "Return the 10 states in Nigeria with their estimated population in JSON format: [{"state": string, "estimated_population": number}]."

**Model Output:**

```json
[
  {"state": "Lagos", "estimated_population": 21700000},
  {"state": "Kano", "estimated_population": 15100000},
  {"state": "Kaduna", "estimated_population": 9310000},
  {"state": "Rivers", "estimated_population": 7600000},
  {"state": "Oyo", "estimated_population": 7400000},
  {"state": "Bauchi", "estimated_population": 6770000},
  {"state": "Katsina", "estimated_population": 6530000},
  {"state": "Anambra", "estimated_population": 5660000},
  {"state": "Kebbi", "estimated_population": 5640000},
  {"state": "Delta", "estimated_population": 5230000}
]
```

✅ JSON is valid — 10 records, consistent structure, correct data types throughout.

---

## Step 4: Comparison Analysis

The key difference between the unstructured and structured prompts lies in format, organization, and machine readability. JSON format matters for real-world applications because it provides a standard, structured way to store and exchange data that is both human readable and machine readable. Testing with multiple inputs across continents, planets, and Nigerian states confirmed that the structured data holds across completely different datasets with different field names and data types, proving the reliability of structured prompts.

---

## Summary Table

| Aspect | Unstructured Prompt | Structured JSON Prompt |
|---|---|---|
| Output Format | Numbered list | Valid JSON array |
| Machine Readable | No | Yes |
| Field Names Defined | No | Yes — enforced by schema |
| Consistent Across Inputs | N/A | Yes — verified across 3 datasets |
| Suitable for API/Program Use | No | Yes |
