package strain

// Ints is an array of ints
type Ints []int

// Lists blah blah
type Lists [][]int

// Strings is an array of strings
type Strings []string

// Keep returns an array of ints where the predicate returns true
func (ints Ints) Keep(predicate func(int) bool) Ints {
	var k Ints
	for _, num := range ints {
		if predicate(num) {
			k = append(k, num)
		}
	}
	return k
}

// Discard returns an array of ints where the predicate returns false
func (ints Ints) Discard(predicate func(int) bool) Ints {
	var d Ints
	for _, num := range ints {
		if !predicate(num) {
			d = append(d, num)
		}
	}
	return d
}

// Keep returns Lists where the predicate returns true
func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var k Lists
	for _, arr := range lists {
		if predicate(arr) {
			k = append(k, arr)
		}
	}
	return k
}

// Keep returns Strings where the predicate returns true
func (strings Strings) Keep(predicate func(string) bool) Strings {
	var k Strings
	for _, s := range strings {
		if predicate(s) {
			k = append(k, s)
		}
	}
	return k
}
