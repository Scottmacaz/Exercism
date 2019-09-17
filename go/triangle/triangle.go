// Package triangle is used to determine triangle types
package triangle

import "math"

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

// KindFromSides determines a type of triangle based on the lengths of its sides.
// If the triangle is not a valid triangle then NaT will be returned.
func KindFromSides(a, b, c float64) Kind {
	sidesEqual := sidesEqual(a, b, c)

	if isTriangle(a, b, c) == false {
		return NaT
	}

	switch sidesEqual {
	case 3:
		return Equ
	case 2:
		return Iso
	case 0:
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	//all must be numbers.
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}

	//If any side is less than or equal to zero then it's not a triangle.
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	//The sum of the lengths of any two sides must be greater than or equal to the
	//length of the third side.
	if a+b < c || a+c < b || b+c < a {
		return false
	}

	return true
}
func sidesEqual(a, b, c float64) int {
	se := 0 // no sides equal.
	if a == b {
		se = 2
		if a == c {
			se++
		}
		return se
	}

	if a == c {
		se = 2
		if b == c {
			se++
		}
		return se
	}

	if b == c {
		se = 2
		if b == a {
			se++
		}
		return se
	}

	return se
}
