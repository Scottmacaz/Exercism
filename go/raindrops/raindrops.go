package raindrops

import "strconv"

//Convert takes a raindrop int and returns the corresponding string
func Convert(r int) string {
	rs := "" //The raindrop string

	if r%3 == 0 {
		rs += "Pling"
	}

	if r%5 == 0 {
		rs += "Plang"
	}

	if r%7 == 0 {
		rs += "Plong"
	}

	if len(rs) == 0 {
		rs = strconv.Itoa(r)
	}
	return rs
}
