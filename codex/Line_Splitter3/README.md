Exercise 3: Line Splitter

You are now simulating how a banner file works.

Input File (sample.txt):
Hello
World
Go


🎯 Your Task

Write a program that:

Reads the file
Splits it into lines
Prints each line with its index

✅ Expected Output
0: Hello
1: World
2: Go

🔒 Requirements
Use:
strings.Split(string(data), "\n")
Handle empty lines properly
Use slice indexing


💡 Hint
lines := strings.Split(string(data), "\n")

for i, line := range lines {
    fmt.Println(i, line)
}