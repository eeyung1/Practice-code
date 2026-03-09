package main

import (
	//"fmt"
	"strings"
	"sort"
)

func IsAnagram(s1, s2 string) bool {
    // convert both strings to lowercase
    s1 = strings.ToLower(s1)
    s2 = strings.ToLower(s2)

    // if lengths differ they can't be anagrams
    if len(s1) != len(s2) {
        return false
    }

    // convert both strings to rune slices
    runes1 := []rune(s1)
    runes2 := []rune(s2)


    // sort both slices
    sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
    sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

    // compare sorted strings
    return string(runes1) == string(runes2)
}

// func main(){
// 	fmt.Println(IsAnagram("word", "rowd"))
// }




// The logic flow in plain English:

// Make both strings lowercase so case doesn't matter
// If lengths are different → immediately false
// Sort both strings alphabetically
// If sorted versions match → they're anagrams