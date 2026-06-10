Extension 1 (10 minutes)
Ignore uppercase and lowercase differences.
Input:
go run . Hello
Output:
h -> 1
e -> 1
l -> 2
o -> 1
Hint:
strings.ToLower()

Extension 2 (15 minutes)
Read the word from a file.
File:
banana
Run:
go run . data.txt
Output:
b -> 1
a -> 3
n -> 2
Hint:
os.ReadFile()
This introduces the file handling they'll need for the banner files in the ASCII Art project.

