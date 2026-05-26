package main

import (
	"reflect"
	"testing"
)

// rows is a shorthand to build a named block of art rows for readability.
func rows(vals ...string) []string { return vals }

func TestStackTwo_CombinesRowsInOrder(t *testing.T) {
	top := rows("a", "b")
	bottom := rows("c", "d")
	got := StackTwo(top, bottom)
	want := rows("a", "b", "c", "d")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStackTwo_TopEmpty(t *testing.T) {
	got := StackTwo(rows(), rows("c", "d"))
	want := rows("c", "d")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStackTwo_BottomEmpty(t *testing.T) {
	got := StackTwo(rows("a", "b"), rows())
	want := rows("a", "b")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStackTwo_BothEmpty(t *testing.T) {
	got := StackTwo(rows(), rows())
	if got == nil {
		t.Error("must not return nil for two empty inputs")
	}
	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestStackTwo_LengthIsSum(t *testing.T) {
	top := rows("a", "b", "c", "d", "e", "f", "g", "h")
	bottom := rows("1", "2", "3", "4", "5", "6", "7", "8")
	got := StackTwo(top, bottom)
	if len(got) != len(top)+len(bottom) {
		t.Errorf("expected length %d, got %d", len(top)+len(bottom), len(got))
	}
}

func TestStackTwo_TopRowsAppearFirst(t *testing.T) {
	top := rows("TOP0", "TOP1", "TOP2", "TOP3", "TOP4", "TOP5", "TOP6", "TOP7")
	bottom := rows("BOT0", "BOT1", "BOT2", "BOT3", "BOT4", "BOT5", "BOT6", "BOT7")
	got := StackTwo(top, bottom)
	for i := 0; i < 8; i++ {
		if got[i] != top[i] {
			t.Errorf("position %d should be top row: got %q, want %q", i, got[i], top[i])
		}
	}
}

func TestStackTwo_BottomRowsAppearAfterTop(t *testing.T) {
	top := rows("T", "T", "T", "T", "T", "T", "T", "T")
	bottom := rows("B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7")
	got := StackTwo(top, bottom)
	for i := 0; i < 8; i++ {
		if got[8+i] != bottom[i] {
			t.Errorf("position %d should be bottom row %d: got %q, want %q",
				8+i, i, got[8+i], bottom[i])
		}
	}
}

func TestStackTwo_DoesNotMutateInputs(t *testing.T) {
	top := rows("a", "b")
	bottom := rows("c", "d")
	origTop := rows("a", "b")
	origBot := rows("c", "d")
	result := StackTwo(top, bottom)
	result[0] = "MUTATED"
	if !reflect.DeepEqual(top, origTop) {
		t.Error("top slice was mutated by StackTwo")
	}
	if !reflect.DeepEqual(bottom, origBot) {
		t.Error("bottom slice was mutated by StackTwo")
	}
}
