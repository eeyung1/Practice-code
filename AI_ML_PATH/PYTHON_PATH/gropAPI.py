from groq import Groq
import os

client = Groq(
    api_key=os.getenv("GROQ_API_KEY")
)

prompt = input("Ask Groq something: ")

response = client.chat.completions.create(
    model="llama-3.3-70b-versatile",
    messages=[
        {
            "role": "user",
            "content": prompt
        }
    ]
)

reply = response.choices[0].message.content


print(f"\nGroq says:\n{reply}")

# print("Groq says:")
# print(reply)

# print("\nSecond print:")
# print(reply)
