// Package acronym contains a method for turning a word into an acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a word and returns the corresponding acronym
func Abbreviate(s string) string {

	sSplit := regexp.MustCompile("[\\-\\ \\_\\s]+").Split(s, -1)
	acronym := ""
	for _, v := range sSplit {
		acronym += strings.ToUpper(string(v[0]))
	}
	return acronym
}
