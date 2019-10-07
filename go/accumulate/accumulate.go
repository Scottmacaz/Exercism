package accumulate

// Accumulate returns a new collection after applying the function defined as f to it
func Accumulate(s []string, f func(c string) string) []string {
	sOut := make([]string, len (s))
	for i, item := range s {
		sOut[i] = f(item)
	}
	return sOut
}
