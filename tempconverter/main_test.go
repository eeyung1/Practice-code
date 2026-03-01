package main

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	result := celsiusToFahrenheit(100) // pass 100 directly
	if result != 212 {                 // we know 100°C = 212°F
		t.Errorf("got %v, want 212", result)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	result := fahrenheitToCelsius(32) // pass 32 directly
	if result != 0 {                  // we know 32°F = 0°C
		t.Errorf("got %v, want 0", result)
	}
}

func TestNegativeTemperature(t *testing.T) {
	result := celsiusToFahrenheit(-40)
	if result != -40 {
		t.Errorf("got %v, want -40", result)
	}
}
