package main

import "testing"

func TestAddition(t *testing.T) {
	result, err := calculate(5, "+", 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("got %v, want 8", result)
	}
}

func TestSubtraction(t *testing.T) {
	result, err := calculate(7, "-", 4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 3 {
		t.Errorf("got %v, want 3", result)
	}
}

func TestMultiplication(t *testing.T) {
	result, err := calculate(3, "*", 6)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 18 {
		t.Errorf("got %v, want 18", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := calculate(10, "/", 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("got %v, want 5", result)
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := calculate(10, "/", 0)
	if err == nil {
		t.Errorf("expected error for division by zero, got nil")
	}
}

func TestUnknownOperator(t *testing.T) {
	_, err := calculate(5, "@", 3)
	if err == nil {
		t.Errorf("expected error for unknown operator, got nil")
	}
}

func TestDecimals(t *testing.T) {
	result, err := calculate(5.5, "+", 3.2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 8.7 {
		t.Errorf("got %v, want 8.7", result)
	}
}