A slice like this:

lines := []string{
	"",
	"A1", "A2", "A3",
	"",
	"B1", "B2", "B3",
}

🎯 Task

Write a function:

func getCharLines(lines []string, index int) []string


Behavior
Each character takes 4 lines total:
1 empty separator
3 actual lines

So:

index = 0 → returns A
index = 1 → returns B

✅ Expected:
getCharLines(lines, 0)
// ["A1", "A2", "A3"]

getCharLines(lines, 1)
// ["B1", "B2", "B3"]


💡 Hint (Core Formula)
start := index * 4 + 1
end   := start + 3