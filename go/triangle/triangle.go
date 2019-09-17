// Package triangle is used to determine triangle types
package triangle


// Kind is a type that represents the kind of triangle
type Kind int

const (
	// NaT represents "not a triangle"
	NaT Kind = iota 
	// Equ represents an equilateral triangle
		Equ 
		// Iso represents an isosceles triangle
		Iso 
		// Sca represents a scalene triangle
    Sca 
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k
}
