// Package bob is a package for responding to Bob
package bob

import (
	"strings"
	"unicode"
)

// Hey returns a canned response from Bob given a type of remark.
func Hey(remark string) string {

	//trim spaces
	trimmed := strings.TrimSpace(remark)

	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}
	punctuation := trimmed[len(trimmed)-1:]

	switch punctuation {
	case "?":
		if isYelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	default:
		if isYelling(remark) {
			return "Whoa, chill out!"
		}
		return "Whatever."
	}
}

func isYelling(remark string) bool {
	if remark == strings.ToUpper(remark) && containsCharacter(remark) {
		return true
	}
	return false
}

func containsCharacter(remark string) bool {
	for _, v := range remark {
		if unicode.IsLetter(v){
			return true
		}
	}
	return false
}

