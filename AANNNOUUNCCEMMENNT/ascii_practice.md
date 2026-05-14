Question 2

Write a function that joins two banner maps into one


Given two banner maps, return a new map that contains all entries from both. If a rune appears in both maps, the entry from priority wins over the entry from base. Neither input map should be modified. This is useful for creating a custom banner that overrides only certain characters.

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string

examples

MergeBanners(base, empty) → copy of base, unmodified

MergeBanners(empty, priority) → copy of priority

if both have 'A', result['A'] must equal priority['A'], not base['A']

runes only in base must still appear in the result

modifying the result map must not affect base or priority

The returned map must be a new allocation — not a reference to either input. Modifying it after the call must not affect base or priority.


Allocate result := make(map[rune][]string). Copy all entries from base, then copy all entries from priority (overwriting base entries for duplicate keys). Return result.



ascii/

├── main.go

├── validate_banner.go

├── merge_banners.go

│

├── validate_banner_test.go

├── merge_banners_test.go

│

├── go.mod

└── README.md




Question 3

Write a function that trims trailing spaces from every row of rendered art


After rendering a line of ASCII art, each row may end with trailing spaces that came from characters like space or certain punctuation. Write a function that takes a slice of 8 art rows and returns a new slice of 8 rows with all trailing spaces removed from each row. The row structure and count must stay the same — only trailing spaces are trimmed.

func TrimArtRows(rows []string) []string

examples

TrimArtRows([]string{"_ ", "| |", ...}) → []string{"_", "| |", ...} (trailing spaces removed per row)

TrimArtRows([]string{"", "", "", "", "", "", "", ""}) → 8 empty strings, unchanged

TrimArtRows([]string{" hi ", ...}) → []string{" hi", ...} (leading spaces preserved)

len of result must always equal len of input

Only trailing spaces must be removed — leading spaces are part of the art and must be preserved. Do not modify the input slice in place; return a new slice.


Allocate result := make([]string, len(rows)). For each row, use strings.TrimRight(row, " ") to strip trailing spaces only. Return result.


ascii/

├── main.go

├── validate_banner.go

├── merge_banners.go

├── trim_art_rows.go

│

├── validate_banner_test.go

├── merge_banners_test.go

├── trim_art_rows_test.go

│

├── go.mod

└── README.md








package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(banner))
	}
	for i, v := range banner {
		if i < 32 && i > 126 {
			return fmt.Errorf("missing character '%c' (Ascii: %c)", i, i)
		}
		if len(v) != 8 {
			return fmt.Errorf("character %c has %d lines, expected 8", i, len(v))
		}
	}
	return nil
}
package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	res := make(map[rune][]string)
	for i, v := range base {
		res[i] = v
	}
	for i, v := range priority {
		res[i] = v
	}
	return res
}
package main

import "strings"

func TrimArtRows(rows []string) []string {
	res := make([]string, len(rows))
	for i, v := range rows {
		if strings.Trim(v, " ") == " " {
			res[i] = v
			continue
		}
		res[i] = strings.TrimRight(v, " ")
	}
	return res
}
package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	
	res := make([]string, len(rows))
	for i, v := range rows {

		if width <= 0 {
		res[i] = v
		}
		padding := width - len(v)
		if padding > 0 {
			res[i] = v + strings.Repeat(" ", padding)
		} else {
			res[i] = v
		}
	}
	return res
}
package main

func StackTwo(top []string, bottom []string) []string {
	res := make([]string, len(top)+len(bottom))
	copy(res, top)
	copy(res[len(top):], bottom)

	return res
}

func StackAll(blocks [][]string) []string {
	res := []string{}
	for _, v := range blocks {
		res = StackTwo(res, v)
	}
	return res
}















package main

import (
	"reflect"
	"strings"
	"testing"
)

// makeArt builds a simple 8-line art slice with an identifying label.
func makeArt(label string) []string {
	art := make([]string, 8)
	for i := range art {
		art[i] = label
	}
	return art
}

func TestMergeBanners_BaseOnlyEntries(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['A'], makeArt("base-A")) {
		t.Errorf("entry only in base must appear in result")
	}
}

func TestMergeBanners_PriorityOnlyEntries(t *testing.T) {
	base := map[rune][]string{}
	priority := map[rune][]string{'B': makeArt("priority-B")}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['B'], makeArt("priority-B")) {
		t.Errorf("entry only in priority must appear in result")
	}
}

func TestMergeBanners_PriorityWinsOnConflict(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	result := MergeBanners(base, priority)
	if !reflect.DeepEqual(result['A'], makeArt("priority-A")) {
		t.Errorf("priority entry must overwrite base entry for same rune")
	}
}

func TestMergeBanners_BothContributeDistinctKeys(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'B': makeArt("priority-B")}
	result := MergeBanners(base, priority)
	if _, ok := result['A']; !ok {
		t.Error("'A' from base is missing in result")
	}
	if _, ok := result['B']; !ok {
		t.Error("'B' from priority is missing in result")
	}
}

func TestMergeBanners_DoesNotModifyBase(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	MergeBanners(base, priority)
	if !reflect.DeepEqual(base['A'], makeArt("base-A")) {
		t.Error("MergeBanners must not modify the base map")
	}
}

func TestMergeBanners_DoesNotModifyPriority(t *testing.T) {
	base := map[rune][]string{}
	priority := map[rune][]string{'A': makeArt("priority-A")}
	MergeBanners(base, priority)
	if !reflect.DeepEqual(priority['A'], makeArt("priority-A")) {
		t.Error("MergeBanners must not modify the priority map")
	}
}

func TestMergeBanners_ResultIsNewMap(t *testing.T) {
	base := map[rune][]string{'A': makeArt("base-A")}
	priority := map[rune][]string{}
	result := MergeBanners(base, priority)
	// Mutating result must not affect base
	result['A'] = makeArt("mutated")
	if reflect.DeepEqual(base['A'], makeArt("mutated")) {
		t.Error("mutating the result must not affect the base map")
	}
}

func TestMergeBanners_BothEmpty(t *testing.T) {
	result := MergeBanners(map[rune][]string{}, map[rune][]string{})
	if result == nil {
		t.Error("result must not be nil even when both inputs are empty")
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d entries", len(result))
	}
}

func TestMergeBanners_ResultLength(t *testing.T) {
	base := map[rune][]string{
		'A': makeArt("A"),
		'B': makeArt("B"),
	}
	priority := map[rune][]string{
		'B': makeArt("B-override"),
		'C': makeArt("C"),
	}
	result := MergeBanners(base, priority)
	// Unique keys: A, B, C = 3
	if len(result) != 3 {
		t.Errorf("expected 3 entries in merged result, got %d", len(result))
	}
}

func TestMergeBanners_NilBaseActsAsEmpty(t *testing.T) {
	priority := map[rune][]string{'A': makeArt("A")}
	// Should not panic on nil base
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MergeBanners panicked on nil base: %v", r)
		}
	}()
	result := MergeBanners(nil, priority)
	if result == nil {
		t.Error("result must not be nil")
	}
}

func TestPadArtRows_PadsShortRows(t *testing.T) {
	input := []string{"hi", "there", "a", "b", "c", "d", "e", "f"}
	result := PadArtRows(input, 8)
	for i, row := range result {
		if len(row) != 8 {
			t.Errorf("row %d: expected length 8, got %d (%q)", i, len(row), row)
		}
	}
}

func TestPadArtRows_PaddingIsSpaces(t *testing.T) {
	input := []string{"ab", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)
	if result[0] != "ab   " {
		t.Errorf("expected \"ab   \", got %q", result[0])
	}
	// All padding characters must be spaces
	if strings.TrimLeft(result[0][2:], " ") != "" {
		t.Errorf("padding must be spaces only, got %q", result[0][2:])
	}
}

func TestPadArtRows_DoesNotTruncate(t *testing.T) {
	// Row wider than target width must come back unchanged
	input := []string{"hello world", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)
	if result[0] != "hello world" {
		t.Errorf("wider row must not be truncated: got %q", result[0])
	}
}

func TestPadArtRows_ExactWidthUnchanged(t *testing.T) {
	input := []string{"abcd", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)
	if result[0] != "abcd" {
		t.Errorf("row at exact width must be unchanged: got %q", result[0])
	}
}

func TestPadArtRows_EmptyRowPaddedToWidth(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)
	for i, row := range result {
		if row != "    " {
			t.Errorf("row %d: empty row should pad to 4 spaces, got %q", i, row)
		}
	}
}

func TestPadArtRows_LeadingSpacesPreserved(t *testing.T) {
	// Leading spaces are art — must not be touched
	input := []string{" _ ", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 8)
	if !strings.HasPrefix(result[0], " _ ") {
		t.Errorf("leading spaces removed: got %q", result[0])
	}
}

func TestPadArtRows_LengthAlwaysPreserved(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh"}
	result := PadArtRows(input, 6)
	if len(result) != len(input) {
		t.Errorf("number of rows changed: got %d, want %d", len(result), len(input))
	}
}

func TestPadArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi", "a", "b", "c", "d", "e", "f", "g"}
	originals := make([]string, len(input))
	copy(originals, input)
	PadArtRows(input, 10)
	for i := range input {
		if input[i] != originals[i] {
			t.Errorf("row %d: input was mutated — must not modify input slice", i)
		}
	}
}

func TestPadArtRows_ZeroWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PadArtRows panicked with width=0: %v", r)
		}
	}()
	input := []string{"hi", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 0)
	if len(result) != 8 {
		t.Errorf("expected 8 rows returned, got %d", len(result))
	}
}

func TestPadArtRows_NegativeWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PadArtRows panicked with width=-5: %v", r)
		}
	}()
	input := []string{"hi", "", "", "", "", "", "", ""}
	PadArtRows(input, -5)
}

func TestPadArtRows_AllRowsSameWidthAfterPadding(t *testing.T) {
	input := []string{"a", "bb", "ccc", "d", "ee", "f", "ggg", "hh"}
	result := PadArtRows(input, 10)
	for i, row := range result {
		if len(row) != 10 {
			t.Errorf("row %d: expected length 10 after padding, got %d (%q)", i, len(row), row)
		}
	}
}

func TestTrimArtRows_RemovesTrailingSpaces(t *testing.T) {
	input := []string{"hello ", "world ", "foo", "bar ", "", " ", "a ", " b "}
	result := TrimArtRows(input)
	want := []string{"hello", "world", "foo", "bar", "", "", "a", " b"}
	for i := range want {
		if result[i] != want[i] {
			t.Errorf("row %d: got %q, want %q", i, result[i], want[i])
		}
	}
}

// func TestTrimArtRows_PreservesLeadingSpaces(t *testing.T) {
// input := []string{" art ", " art ", " a", "a ", " ", " ", "x", ""}
// result := TrimArtRows(input)
// for i, row := range result {
// if strings.HasPrefix(input[i], " ") && !strings.HasPrefix(row, " ") {
// t.Errorf("row %d: leading spaces removed — must be preserved. got %q", i, row)
// }
// }
// }

func TestTrimArtRows_LengthUnchanged(t *testing.T) {
	input := []string{"a ", "b", "c ", "d", "e ", "f", "g ", "h"}
	result := TrimArtRows(input)
	if len(result) != len(input) {
		t.Errorf("expected %d rows, got %d", len(input), len(result))
	}
}

func TestTrimArtRows_AllEmptyRows(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := TrimArtRows(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: expected empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_AllSpaceRows(t *testing.T) {
	input := []string{" ", " ", " ", " ", " ", " ", " ", " "}
	result := TrimArtRows(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: all-space row should become empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_NoTrailingSpaces(t *testing.T) {
	// Nothing to trim — rows should come back identical
	input := []string{"_", "| |", "|_|", "", " _", "| |", "|_|", ""}
	result := TrimArtRows(input)
	for i := range input {
		if result[i] != input[i] {
			t.Errorf("row %d: no trailing spaces, must be unchanged. got %q, want %q",
				i, result[i], input[i])
		}
	}
}

func TestTrimArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi ", "there ", "a ", "b ", "c ", "d ", "e ", "f "}
	originals := make([]string, len(input))
	copy(originals, input)
	TrimArtRows(input)
	for i := range input {
		if input[i] != originals[i] {
			t.Errorf("row %d: input was modified — must return new slice. got %q", i, input[i])
		}
	}
}

func TestTrimArtRows_ReturnsNewSlice(t *testing.T) {
	input := []string{"a ", "b ", "c ", "d ", "e ", "f ", "g ", "h "}
	result := TrimArtRows(input)
	// Mutate result — original must be unchanged
	result[0] = "MUTATED"
	if input[0] == "MUTATED" {
		t.Error("TrimArtRows must return a new slice, not a reference to the input")
	}
}

func TestTrimArtRows_MidRowSpacesPreserved(t *testing.T) {
	// Internal spaces in art must never be touched
	input := []string{"| _ | ", "| | | ", "|___|", "", "", "", "", ""}
	result := TrimArtRows(input)
	if !strings.Contains(result[0], "| _ |") {
		t.Errorf("row 0: internal spaces removed — must be preserved. got %q", result[0])
	}
	if !strings.Contains(result[1], "| | |") {
		t.Errorf("row 1: internal spaces removed — must be preserved. got %q", result[1])
	}
}

func TestTrimArtRows_EmptySlice(t *testing.T) {
	result := TrimArtRows([]string{})
	if result == nil {
		t.Error("must not return nil for empty input slice")
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got length %d", len(result))
	}
}

// buildGoodBanner constructs a perfectly valid banner map in memory.
// No file needed — all 95 printable ASCII chars, each with exactly 8 lines.
func buildGoodBanner() map[rune][]string {
	banner := make(map[rune][]string)
	for r := rune(32); r <= 126; r++ {
		art := make([]string, 8)
		for i := range art {
			art[i] = string(r) + "row"
		}
		banner[r] = art
	}
	return banner
}

func TestValidateBanner_GoodMap(t *testing.T) {
	err := ValidateBanner(buildGoodBanner())
	if err != nil {
		t.Errorf("expected nil for valid banner, got: %v", err)
	}
}

func TestValidateBanner_Nil(t *testing.T) {
	err := ValidateBanner(nil)
	if err == nil {
		t.Error("expected error for nil banner, got nil")
	}
}

func TestValidateBanner_WrongEntryCount(t *testing.T) {
	banner := buildGoodBanner()
	delete(banner, 'A')
	delete(banner, 'B')
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for banner with 93 entries, got nil")
	}
}

func TestValidateBanner_MissingSpace(t *testing.T) {
	banner := buildGoodBanner()
	delete(banner, ' ')
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for missing space character")
	}
}

func TestValidateBanner_MissingTilde(t *testing.T) {
	// ~ is ASCII 126 — highest valid character
	banner := buildGoodBanner()
	delete(banner, '~')
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for missing tilde (~) character")
	}
}

func TestValidateBanner_TooFewLines(t *testing.T) {
	banner := buildGoodBanner()
	banner['A'] = []string{"only", "six", "lines", "here", "not", "eight"}
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error when character 'A' has 6 lines instead of 8")
	}
	if err != nil && !strings.Contains(err.Error(), "A") {
		t.Errorf("error message should mention the offending character 'A', got: %v", err)
	}
}

func TestValidateBanner_TooManyLines(t *testing.T) {
	banner := buildGoodBanner()
	banner['Z'] = make([]string, 10)
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error when character 'Z' has 10 lines instead of 8")
	}
}

func TestValidateBanner_ZeroLinesForChar(t *testing.T) {
	banner := buildGoodBanner()
	banner['!'] = []string{}
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error when character '!' has 0 lines")
	}
}

func TestValidateBanner_EmptyMap(t *testing.T) {
	err := ValidateBanner(map[rune][]string{})
	if err == nil {
		t.Error("expected error for empty map, got nil")
	}
}

func TestValidateBanner_ExtraCharacterDoesNotMakeItValid(t *testing.T) {
	// A map with 95 correct entries + 1 extra non-ASCII rune still has the wrong count
	banner := buildGoodBanner()
	banner[rune(200)] = make([]string, 8) // add a non-printable-ASCII char
	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error: 96 entries should fail the count check")
	}
}

func TestValidateBanner_ErrorMessageIsDescriptive(t *testing.T) {
	banner := buildGoodBanner()
	banner['M'] = []string{"only", "three"}
	err := ValidateBanner(banner)
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if len(err.Error()) < 5 {
		t.Errorf("error message too short to be descriptive: %q", err.Error())
	}
}

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

func TestStackAll_EmptyInput(t *testing.T) {
	got := StackAll([][]string{})
	if got == nil {
		t.Error("StackAll must not return nil for empty input")
	}
	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestStackAll_SingleBlock(t *testing.T) {
	block := rows("a", "b", "c", "d", "e", "f", "g", "h")
	got := StackAll([][]string{block})
	if !reflect.DeepEqual(got, block) {
		t.Errorf("StackAll with one block must equal that block: got %v", got)
	}
}

func TestStackAll_ThreeBlocks(t *testing.T) {
	a := rows("A1", "A2")
	b := rows("B1", "B2")
	c := rows("C1", "C2")
	got := StackAll([][]string{a, b, c})
	want := rows("A1", "A2", "B1", "B2", "C1", "C2")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStackAll_TotalLengthIsSum(t *testing.T) {
	blocks := [][]string{
		rows("a", "b", "c", "d", "e", "f", "g", "h"),
		rows("1", "2", "3", "4", "5", "6", "7", "8"),
		rows("x", "y", "z", "w", "v", "u", "t", "s"),
	}
	got := StackAll(blocks)
	if len(got) != 24 {
		t.Errorf("expected 24 rows (3 × 8), got %d", len(got))
	}
}

func TestStackAll_OrderPreserved(t *testing.T) {
	blocks := [][]string{
		rows("FIRST"),
		rows("SECOND"),
		rows("THIRD"),
	}
	got := StackAll(blocks)
	if got[0] != "FIRST" || got[1] != "SECOND" || got[2] != "THIRD" {
		t.Errorf("block order not preserved: got %v", got)
	}
}

func TestStackAll_NilInput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("StackAll panicked on nil input: %v", r)
		}
	}()
	got := StackAll(nil)
	if got == nil {
		t.Error("must return non-nil slice even for nil input")
	}
}

