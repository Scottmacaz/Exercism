// Package proverb creates proverbs from a list of nouns
package proverb

import (
	"fmt"
)

// Proverb will produce a proverb in the form of: "For want of a [noun] the [noun] was lost".
// Each subsequent line will use the next noun from the list.
func Proverb(rhyme []string) []string {

	proverb := []string{}

	for i := 0; i < len(rhyme); i++ {
		if i != len(rhyme)-1 {
			line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
			proverb = append(proverb, line)
		} else {
			lastLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			proverb = append(proverb, lastLine)
		}
	}
	return proverb
}
