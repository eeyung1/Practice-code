# ASCII-Art Learning Path

## Quest 1: Character Counter

Goal:
Count how many times each character appears.

Input:
go run . banana

Output:
b -> 1
a -> 3
n -> 2

Learn:

* os.Args
* Maps
* Loops
* Error handling

Connection:
ASCII-Art eventually needs to process characters individually.

---

## Quest 2: Case Insensitivity

Goal:
Treat uppercase and lowercase as the same.

Input:
go run . Hello

Output:
h -> 1
e -> 1
l -> 2
o -> 1

Learn:

* strings.ToLower()

Connection:
Introduces string manipulation.

---

## Quest 3: Reading From Files

Goal:
Read input from a text file.

File:
banana

Run:
go run . data.txt

Output:
b -> 1
a -> 3
n -> 2

Learn:

* os.ReadFile()

Connection:
Banner files are read exactly the same way.

---

## Quest 4: Character Position

Goal:
Print the ASCII code of each character.

Input:
go run . Hi

Output:
H -> 72
i -> 105

Learn:

* Rune iteration
* int(ch)

Connection:
ASCII-Art computes positions from character codes.

---

## Quest 5: Character Index Mapping

Goal:
Map characters to positions.

Input:
go run . A

Output:
A -> 33

Learn:

* int(ch)-32

Connection:
The final project locates each character inside the banner file using this formula.

---

## Quest 6: Slices

Goal:
Split text into lines.

Input:
Hello
There

Output:
[Hello There]

Learn:

* strings.Split()

Connection:
Needed for handling \n.

---

## Quest 7: Nested Loops

Goal:
Print a rectangle.

Input:
3 5

Output:

---

---

---

Learn:

* Nested loops

Connection:
ASCII-Art uses nested loops extensively.

---

## Quest 8: Multi-Line Characters

Goal:
Represent each character using multiple lines.

Example:

A:

---

* *

---

Learn:

* Slices of strings
* Nested loops

Connection:
ASCII-Art characters have height 8.

---

## Quest 9: Mini Banner

Goal:
Create a banner containing only A and B.

Input:
AB

Output:
(large representation)

Learn:

* Character lookup
* Indexing

Connection:
Introduces the idea behind banner files.

---

## Quest 10: Parsing Banner Files

Goal:
Read a banner file and split it into lines.

Learn:

* strings.Split()
* File parsing

Connection:
Exactly what the final project does.

---

## Quest 11: Character Height

Goal:
Understand that every character occupies 8 lines.

Learn:

* Fixed-size blocks

Connection:
Characters in the banner file have height 8.

---

## Quest 12: Character Separation

Goal:
Understand why each character occupies 9 lines.

Learn:

* Empty separator lines

Connection:
Index formula:

(int(ch)-32)*9

---

## Quest 13: Rendering One Character

Goal:
Print a large A.

Learn:

* Index calculations

Connection:
Core rendering logic.

---

## Quest 14: Rendering One Word

Input:
HELLO

Learn:

* Outer loop over rows
* Inner loop over characters

Connection:
The fundamental algorithm.

---

## Quest 15: Supporting Spaces

Input:
HELLO THERE

Learn:

* Space character handling

Connection:
ASCII-Art supports spaces.

---

## Quest 16: Supporting Numbers

Input:
123

Learn:

* Character ranges

Connection:
ASCII-Art supports numbers.

---

## Quest 17: Supporting Special Characters

Input:
{Hello}

Learn:

* Generic character handling

Connection:
Banner files support all printable ASCII.

---

## Quest 18: Error Handling

Goal:
Detect invalid inputs.

Learn:

* Input validation
* Error messages

Connection:
Required for robustness.

---

## Quest 19: Supporting "\n"

Input:
Hello\nThere

Learn:

* strings.Split()

Connection:
Required by project.

---

## Quest 20: Supporting Consecutive Newlines

Input:
Hello\n\nThere

Learn:

* Empty strings
* Edge cases

Connection:
One of the trickiest parts of the project.

---

## Quest 21: Multiple Banner Files

Goal:
Allow:

standard
shadow
thinkertoy

Learn:

* Dynamic file paths

Connection:
Required by bonus projects.

---

## Quest 22: Refactoring

Goal:
Split logic into functions.

Functions:

* ReadBanner()
* GenerateAscii()
* ValidateInput()

Learn:

* Separation of concerns

Connection:
Professional code organization.

---

## Quest 23: Unit Tests

Goal:
Test small functions independently.

Learn:

* testing package

Connection:
Good engineering practices.

---

## Quest 24: Full ASCII-Art

Combine:

Arguments
+
Maps
+
Files
+
Strings
+
Slices
+
Nested loops
+
Index calculations
+
Newline handling
+
Functions
+
Error handling

Result:

ASCII-Art
