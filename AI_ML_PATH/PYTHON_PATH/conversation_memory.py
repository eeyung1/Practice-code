from groq import Groq
import os

client = Groq(
    api_key=os.getenv("GROQ_API_KEY")
)

print(client)

messages = []

while True:
    prompt = input("Ask Groq something: ")

    if prompt.lower() == "exit":
        print("Goodbye!")
        break
    
    messages.append(
        {
            "role":"user",
            "content": prompt
        }
    )
    
    response = client.chat.completions.create(
    model="llama-3.3-70b-versatile",
    messages=messages
    )

    reply = response.choices[0].message.content


    print(f"\nGroq says:\n{reply}")
