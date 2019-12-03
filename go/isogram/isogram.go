package isogram

import (
	"unicode"
)

//IsIsogram returns true if a word is an isogram
func IsIsogram(word string) bool {
	charMap := make(map[rune]rune)

	for _, v := range word {
		if string(v) == " " || string(v) == "-" {
			continue
		}

		v = unicode.ToLower(v)
		_, ok := charMap[v]

		if ok {
			return false
		}
		charMap[v] = v
	}
	return true
}
