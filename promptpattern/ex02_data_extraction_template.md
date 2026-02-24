
# Exercise 2: Data Extraction Template

## Objective
Extract structured data from unstructured text using a reusable prompt template and ensure consistent JSON output.

---

## Reusable Extraction Template

Extract the following fields from the text: Name, Age, Location, Occupation, Company.  
If any field is missing, return null.  
Return the result strictly in valid JSON format using the following structure:  
{ "Name": "", "Age": "", "Location": "", "Occupation": "", "Company": "" }  
Do not include any explanation, only valid JSON.

---

## Input 1

Emmanuel Eyung, age 30, lives in Abuja and works as a Software Engineer at Moniepoint.

### Model Output

```json
{
  "Name": "Emmanuel Eyung",
  "Age": 30,
  "Location": "Abuja",
  "Occupation": "Software Engineer",
  "Company": "Moniepoint"
}
```

---

## Input 2

Victor Smith is 35 years old, based in Paris. He is a Brand Manager of Coca cola.

### Model Output

```json
{
  "Name": "Victor Smith",
  "Age": 35,
  "Location": "Paris",
  "Occupation": "Brand Manager",
  "Company": "Coca cola"
}
```

---

## Analysis

The outputs followed the exact JSON structure specified in the prompt. All required keys were present and formatted correctly without additional explanation. The formatting remained consistent across both inputs even though the sentences were structured differently.

Specifying “return strictly JSON” improved reliability by preventing the model from adding extra commentary. This ensures predictable outputs that can be used in automated systems or APIs. Handling missing fields is important because real-world data may be incomplete, and returning null values keeps the structure consistent.

This demonstrates reusable prompt design because the same structured template was applied to different inputs and still produced consistent results.
