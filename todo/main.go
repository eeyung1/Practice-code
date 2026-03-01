package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . add/list/remove")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add 'task'")
			os.Exit(1)
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . remove 1")
			os.Exit(1)
		}
		removeTask(os.Args[2])
	default:
		fmt.Println("Unknown command. Use add, list or remove")
	}
}

func addTask(task string) {
	// step 1: open todos.txt in append mode
	file, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
	}
	defer file.Close()
	// step 2: write the task + newline
	file.WriteString(task + "\n")
	// step 3: confirm to user
	fmt.Println("Added:", task)
}

func listTasks() {
	// step 1: read todos.txt
	data, err := os.ReadFile("todos.txt")
	if err != nil {
		fmt.Println("No tasks found!")
		return
	}
	// step 2: split into lines
	tasks := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	// step 3: loop and print with numbers
	for i, task := range tasks {
		fmt.Println(i+1, ".", task)
	}
}

func removeTask(arg string) {
	data, err := os.ReadFile("todos.txt")
	if err != nil {
		fmt.Println("No tasks found!")
		return
	}
	tasks := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Error: please provide a valid number")
		os.Exit(1)
	}
	if num < 1 || num > len(tasks) {
		fmt.Println("Error: task number out of range")
		os.Exit(1)
	}
	tasks = append(tasks[:num-1], tasks[num:]...)
	os.WriteFile("todos.txt", []byte(strings.Join(tasks, "\n")+"\n"), 0644)
	fmt.Println("Removed task", num)
}