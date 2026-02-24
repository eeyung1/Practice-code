# Exercise 2: Adding Source or Data Constraints

## Objective
Force responses to be tied to given material by adding citation and passage constraints to prompts.

---

## Step 1: Unconstrained Prompt (No Source or Citation)

**Prompt:**
> "How could quantum computing break encryption?"

**Model Output:**

**Why today's encryption is secure (classically)**
Modern encryption relies on math problems that are extremely hard to reverse without a key. Examples include factoring large numbers (RSA) and discrete logarithms (ECC). Classical computers would take impractically long to break strong encryption.

**What quantum computers do differently**
Quantum computers use qubits, superposition, and entanglement to process many possibilities at once. This enables special algorithms that dramatically speed up solving certain problems.

**Shor's Algorithm — breaks public-key encryption**
Developed by Peter Shor. Efficiently factors large numbers and solves discrete logarithms. Could break RSA, ECC, and Diffie-Hellman much faster than classical computers.

**Grover's Algorithm — weakens symmetric encryption**
Speeds up brute-force key searches. Effectively halves key strength (e.g., AES-256 ≈ AES-128 security). Symmetric encryption remains relatively strong if key sizes are increased.

**Why the internet isn't broken yet**
Current quantum computers lack enough stable qubits and low error rates. Breaking real-world encryption would require far more advanced machines than exist today.

**Post-quantum cryptography**
New algorithms are being developed to resist quantum attacks, led by organizations like NIST. Includes lattice-based, hash-based, and code-based cryptography.

---

## Step 2A: Citation-Constrained Prompt

**Prompt:**
> "How quantum computing could break encryption. Cite this source (https://medium.com/@akitrablog/the-invisible-threat-how-quantum-computing-could-break-todays-encryption-888e3ea99cf3) in your answer."

**Model Output:**

**Why encryption works today**
Modern encryption protects data by using mathematical problems that are extremely hard for classical computers to solve, keeping information private across networks and devices. (Medium)
Encryption is used everywhere — from securing communications to verifying identity — and relies on keys known only to authorized parties. (Medium)

**What makes quantum computers different**
Quantum computers use qubits that can exist in multiple states at once (superposition) and be linked together (entanglement), allowing them to process many possibilities simultaneously and solve certain problems much faster than classical computers. (Medium)

**Shor's Algorithm**
Can factor large numbers efficiently. This threatens asymmetric cryptography such as RSA, Diffie-Hellman, and elliptic curve systems. Tasks that would take classical computers trillions of years could theoretically be done much faster on a large quantum computer. (Medium)

**Grover's Algorithm**
Speeds up searching for keys. Mainly affects symmetric encryption by requiring larger keys for equivalent security. (Medium)

**Why asymmetric encryption is most vulnerable**
Asymmetric systems rely on the difficulty of factoring large numbers or solving discrete logarithms — problems quantum computers are particularly good at solving. (Medium)
Symmetric encryption is less affected because there is currently no known quantum method that fully breaks it, though stronger keys may be needed. (Medium)

**Potential consequences**
If powerful quantum computers become available, sensitive data in sectors like finance, healthcare, government, and technology could be exposed if protected by vulnerable algorithms. (Medium)

**Post-quantum cryptography**
Researchers are developing quantum-resistant algorithms designed to remain secure against both classical and quantum attacks. (Medium)
Efforts include new standards such as hybrid cryptography, larger keys, and quantum key distribution to mitigate risks. (Medium)

**Why this isn't happening yet**
Current quantum computers are still too small and error-prone to break real-world encryption, but advances by major tech companies are driving preparation for a future "Y2Q" milestone. (Medium)

---

## Step 2B: Passage-Constrained Prompt

**Reference Passage Provided:**
> The most recent understanding of prime numbers states that while finding large prime numbers is simple, factoring the product of two huge prime numbers may be quite challenging. Prime factors are so hard to compute that experts estimate it would take trillions of years for a current computer to break typical 2,048-bit RSA encryption. In 1995, scientist Peter Shor demonstrated that actual quantum computations could be performed by correcting errors that would maintain quantum states. Shor presented a novel method of computing prime factors using quantum computers that can reduce computation time from trillions of years to just eight hours. All current asymmetric cryptography implementations, including RSA, Diffie-Hellman, and elliptic curve cryptography, are theoretically breakable by quantum computers because they rely on calculations for discovering prime factors. The development of cryptographic algorithms that are incomprehensible to both quantum and classical computers is referred to as post-quantum cryptography.

**Prompt:**
> "Use only the text provided below to explain how quantum computing could break encryption."

**Model Output:**

**Why current encryption is considered secure**
While finding large prime numbers is simple, factoring the product of two huge prime numbers may be quite challenging. Prime factors are so hard to compute that experts estimate it would take trillions of years for a current computer to break typical 2,048-bit RSA encryption.

**How quantum computing changes this**
In 1995, scientist Peter Shor demonstrated that actual quantum computations could be performed by correcting errors that would maintain quantum states. Shor presented a novel method of computing prime factors using quantum computers that can reduce computation time from trillions of years to just eight hours.

**Why asymmetric encryption is vulnerable**
All current asymmetric cryptography implementations, including RSA, Diffie-Hellman, and elliptic curve cryptography, are theoretically breakable by quantum computers because they rely on calculations for discovering prime factors. Symmetric cryptography is less susceptible, and there is currently no mechanism that would be able to defeat it, although a slightly bigger secret key may be necessary.

**Post-quantum cryptography**
The development of cryptographic algorithms that are incomprehensible to both quantum and classical computers is referred to as post-quantum cryptography. It is anticipated that quantum technology will advance to a point where it will allow computations capable of breaching conventional cryptographic systems, especially asymmetric cryptography.

---

## Step 3: Comparison Analysis

The unconstrained prompt sounds confident and draws entirely from the model's general training knowledge with no source material anchoring it. The citation-constrained prompt had citations appearing throughout which shows the constraint had an effect, however citing "Medium" generically without a specific article link does not make it truly reliable. The passage-constrained prompt uses only the text found in the passage to explain the concept, with the AI staying within boundaries rather than drawing from outside knowledge. This makes the passage-constrained prompt the most effective of the three, because keeping the AI within specific boundaries is more valuable than allowing it to be creative with unverified information.

---

## Summary Table

| Aspect | Unconstrained | Citation-Constrained | Passage-Constrained |
|---|---|---|---|
| Data Source | Model's general training | Generic "Medium" citation | Specific provided passage |
| Source Reliability | None | Weak — no specific URL | High — tied to given text |
| Risk of Hallucination | High | Medium | Low |
| Output Focus | Broad and general | Broad with loose attribution | Strictly within passage |
| Most Effective? | No | Partial | Yes |
