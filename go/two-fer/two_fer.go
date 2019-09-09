// Package twofer is used to implement the exercism rules for sharing.
package twofer

import "fmt"

// ShareWith will share with name if one is passed in, otherwise the pronoun "you" will be used.
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	
	return fmt.Sprintf("One for %v, one for me.", name)
}
