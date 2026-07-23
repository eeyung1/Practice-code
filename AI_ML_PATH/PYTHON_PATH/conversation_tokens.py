from groq import Groq
import os

client = Groq(
    api_key=os.getenv("GROQ_API_KEY")
)

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

    prompt = response.usage.prompt_tokens

    completion = response.usage.completion_tokens

    total = response.usage.total_tokens

    messages.append(
        {

            "role": "assistant",

            "content": reply

        }
    )


    print(f"\nGroq says:\n{reply}")
    print(f"Prompt Tokens:\n{prompt}")
    print(f"Completion Tokens:\n{completion}")
    print(f"Total Tokens:\n{total}")
