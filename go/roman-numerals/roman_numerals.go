package romannumerals

import (
	"strconv"
	"strings"
	"errors"
)

// ToRomanNumeral is a function that will convert a 
// number between 1 and 3000 to a roman numeral.
func ToRomanNumeral(a int) (string, error) {

	if (a <= 0 || a > 3000) {
		return "",  errors.New("Number to convert must be between 1 and 3000")
	}
	places := getPlaces(a)
	rn := ""
	for _, p := range places {
		rnu := makeRomanNumeral(p)
		rn += rnu
	}
	return rn, nil
}
func makeRomanNumeral(arabic int) string {
	rn := ""
	for arabic > 0 {
		switch {
		case arabic >= 1000:
			arabic -= 1000
			rn += "M"
		case arabic >= 900:
			arabic -= 900
			rn += "CM"
		case arabic >= 500:
			arabic -= 500
			rn += "D"
		case arabic >= 400:
			arabic -= 400
			rn += "CD"
		case arabic >= 100:
			arabic -= 100
			rn += "C"
		case arabic >= 90:
			arabic -= 90
			rn += "XC"
		case arabic >= 50:
			arabic -= 50
			rn += "L"
		case arabic >= 40:
			arabic -= 40
			rn += "XL"
		case arabic >= 10:
			arabic -= 10
			rn += "X"
		case arabic >= 9:
			arabic -= 9
			rn += "IX"
		case arabic >= 5:
			arabic -= 5
			rn += "V"
		case arabic >= 4:
			arabic -= 4
			rn += "IV"
		case arabic >= 1:
			arabic--
			rn += "I"
		}
	}
	return rn
}

//This is the code that converts an int to a group of places:
func getPlaces(a int) []int {
	ds := strconv.Itoa(a)
	rsp := make([]int, len(ds))
	pos := 1
	for i, v := range ds {
		c := string(v)
		c += strings.Repeat("0", len(ds)-pos)
		pos++
		rsp[i], _ = strconv.Atoi(c)
	}
	return rsp
}