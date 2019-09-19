package hamming

import (
	"errors"
)

// Distance is a function counting the number of differences between two homologous DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("string lengths must be the same")
	}
	hd := 0 //hamming distance
	for i, ai := range a {
		if uint8(ai) != b[i] {
			hd++
		}
	}
	return hd, nil
}
