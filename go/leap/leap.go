// Package leap is a package for implementing the exercism.io leap exercise in Go
package leap

// IsLeapYear takes in a year and returns true if the year is a leap year and false if it is not
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 {
			return true
		} else {
			if year%400 == 0 {
				return true
			}
		}
	}
	return false
}
