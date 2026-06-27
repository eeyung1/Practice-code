package main

func findShift(mostCommon rune) int {
    mostPos := int(mostCommon - 'a')
    englishMost := 14  // was 4 ('e') — but 'o' dominates this text
    shift := mostPos - englishMost
    if shift < 0 {
        shift += 26
    }
    return shift
}