package hamming

import (
	"errors"
)

// Distance is a function counting the number of differences between two homologous DNA strands
func Distance(a, b string) (int, error) {

	hd := 0 //hamming distance
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("string lengths must be the same")
	}

	for i, ai := range ar {
		if ai != br[i] {
			hd++
		}
	}
	return hd, nil
}
