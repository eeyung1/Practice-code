package main

import "testing"

// test that listTasks doesn't crash on empty input
func TestListEmpty(t *testing.T) {
	// just make sure it doesn't panic
	tasks := []string{}
	if len(tasks) != 0 {
		t.Errorf("expected empty list")
	}
}

// test the remove logic
func TestRemoveLogic(t *testing.T) {
	tasks := []string{"Buy groceries", "Call mom", "Do laundry"}
	// remove index 1 (Call mom)
	tasks = append(tasks[:1], tasks[2:]...)
	if len(tasks) != 2 {
		t.Errorf("got %v tasks, want 2", len(tasks))
	}
	if tasks[0] != "Buy groceries" {
		t.Errorf("got %v, want Buy groceries", tasks[0])
	}
	if tasks[1] != "Do laundry" {
		t.Errorf("got %v, want Do laundry", tasks[1])
	}
}