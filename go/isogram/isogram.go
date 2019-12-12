package isogram

import (
	"unicode"
)

//IsIsogram returns true if a word is an isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, v := range word {
		if string(v) == " " || string(v) == "-" {
			continue
		}
		v = unicode.ToLower(v)
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}
