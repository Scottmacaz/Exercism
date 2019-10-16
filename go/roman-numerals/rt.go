package main

import "fmt"

type romanNumeralRank struct {
	rank   int
	arabic int
	roman  string
}

var romanNumeralRanks = []romanNumeralRank{
	{7, 1000, "M"},
	{6, 500, "D"},
	{5, 100, "C"},
	{4, 50, "L"},
	{3, 10, "X"},
	{2, 5, "V"},
	{1, 1, "I"},
}

func makeRomanNumeral(a int) string {

	fmt.Println("Making roman numeral from: ", a)
	rn := ""
	for a > 0 {
		for _, v := range romanNumeralRanks {
			if a-v.arabic >= 0 {
				fmt.Println("Subtracting ", v.arabic, " from ", a)
				rn += v.roman
				a -= v.arabic
				if needsPromoting(rn) {
					rn = promote(rn, v.rank+1)
				}
			}
		}
	}
	return rn
}

func promote(rn string, rank int) string {
	promote := rn[0 : len(rn)-3]
	promote += promote + "B"
	return promote
}

func needsPromoting(r string) bool {
	if len(r) < 3 {
		return false
	}
	if r[len(r)-3] == r[len(r)-2] && r[len(r)-2] == r[len(r)-1] {
		return true
	}
	return false
}

// //This is the code that converts an int to a group of places:
func getPlaces(a int) []int {
	// 	x := 43928
	i := 10
	var places []int

	for a > i/10 {
		place := a%i - a%(i/10)
		places = append(places, place)
		i *= 10
	}
	return places
}

func main() {
	fmt.Println()
	fmt.Println("Hello World")
	fmt.Println("Done! 10 = ", makeRomanNumeral(10), "\n")
	fmt.Println("Done! 11 = ", makeRomanNumeral(10), "\n")
	fmt.Println("Done! 12 = ", makeRomanNumeral(10), "\n")
	fmt.Println("Done! 13 = ", makeRomanNumeral(10), "\n")
	c := 42
	
	bbbbbb
}
