package main

import (
	"testing"
)

func TestCalculateStats (t *testing.T) {
	grades := []float64{85, 92, 78}
	avg, highest, lowest := calculateStats(grades)

	if avg != 85 {
		t.Errorf("avg: got %v, want 85", avg)
	}

	if highest != 92 {
		t.Errorf("highest: got %v, want 92", highest)
	}

	if lowest != 78 {
		t.Errorf("lowest: got %v, want 78", lowest)
	}
}
