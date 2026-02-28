package main

import "testing"

func TestCountWords(t *testing.T) {
	text := "Hello world this is a test"
	result := countWords(text)
	if result != 6 {
		t.Errorf("got %v, want 6", result)
	}
}

func TestCountLines(t *testing.T) {
	text := "Hello world\nthis is a test\nthird line"
	result := countLines(text)
	if result != 3 {
		t.Errorf("got %v, want 3", result)
	}
}

func TestCountChars(t *testing.T) {
	text := "Hello" // 5 characters, no spaces
	result := countChars(text)
	if result != 5 {
		t.Errorf("got %v, want 5", result)
	}
}

func TestCountCharsWithSpaces(t *testing.T) {
	text := "Hello World" // 11 characters including the space
	result := countChars(text)
	if result != 11 {
		t.Errorf("got %v, want 11", result)
	}
}
