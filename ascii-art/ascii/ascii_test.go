package ascii

import (
	"os"
	"strings"
	"testing"
)

func TestLoadBanner(t *testing.T) {
	os.Chdir("../")

	banner, err := LoadBanner("standard")
	if err != nil {
		t.Fatalf("Expected no error loading banner, got: %v", err)
	}

	if len(banner) < 95 {
		t.Errorf("Expected at least 95 characters in banner, got %d", len(banner))
	}

	for i, char := range banner {
		if len(char) != 8 {
			t.Errorf("Character at index %d has %d lines, expected 8", i, len(char))
		}
	}
}

func TestRenderEmpty(t *testing.T) {
	result, err := Render("", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty output for empty input, got: %q", result)
	}
}

func TestRenderNewlineOnly(t *testing.T) {
	os.Chdir("../")

	result, err := Render("\\n", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != "\n" {
		t.Errorf("Expected single newline for input '\\n', got: %q", result)
	}
}

func TestRenderHello(t *testing.T) {
	os.Chdir("../")

	result, err := Render("Hello", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputLines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(outputLines) != 8 {
		t.Errorf("Expected 8 lines for 'Hello', got %d", len(outputLines))
	}
}

func TestRenderSpecialChars(t *testing.T) {
		os.Chdir("../")

	_, err := Render("{Hello There}", "standard")
	if err != nil {
		t.Fatalf("Unexpected error rendering special characters: %v", err)
	}
}
