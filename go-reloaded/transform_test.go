package main

import "testing"

func TestHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestBin(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestUp(t *testing.T) {
	input := "Ready, set, go (up) !"
	expected := "Ready, set, GO!"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestLow(t *testing.T) {
	input := "I should stop SHOUTING (low)"
	expected := "I should stop shouting"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestCap(t *testing.T) {
	input := "Welcome to the Brooklyn bridge (cap)"
	expected := "Welcome to the Brooklyn Bridge"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestUpWithCount(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestArticle(t *testing.T) {
	input := "There is no greater agony than bearing a untold story inside you."
	expected := "There is no greater agony than bearing an untold story inside you."
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestPunctuation(t *testing.T) {
	input := "I was sitting over there ,and then BAMM !!"
	expected := "I was sitting over there, and then BAMM!!"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestQuotes(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"
	result := transform(input)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}