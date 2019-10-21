package isogram

import "strings"

//IsIsogram returns true if a word is an isogram
func IsIsogram(word string) bool {
	charsUsed := ""
	for _, v := range word {
		if string(v) == " " || string(v) == "-" {
			continue
		}
		if strings.Contains(charsUsed, strings.ToLower(string(v))) {
			return false
		}
		charsUsed += strings.ToLower(string(v))
	}
	return true
}
