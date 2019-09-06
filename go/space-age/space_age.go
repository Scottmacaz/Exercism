package space

import "fmt"

// Planet stores a type of planet
type Planet string

// Age takes an age in seconds and a planet.  It returns an age in years
// on the given planet.
func Age(ageInSeconds float64, planet Planet) float64 {
	//The following floats are percentage of earth year.
	orbitalMap := make(map[Planet]float64)
	orbitalMap["Earth"] = 1.0
	orbitalMap["Mercury"] = 0.2408467
	orbitalMap["Venus"] = 0.61519726
	orbitalMap["Mars"] = 1.8808158
	orbitalMap["Jupiter"] = 11.862615
	orbitalMap["Saturn"] = 29.447498
	orbitalMap["Uranus"] = 84.016846
	orbitalMap["Neptune"] = 164.79132

	if val, ok := orbitalMap[planet]; ok {
		return determineRelativeAge(ageInSeconds, val)
	}
	panic(fmt.Sprintf("Unrecognized planet: %v", planet))
}

func determineRelativeAge(ageInSeconds float64, planetOrbitalPercentage float64) float64 {
	secondsInDay := float64(86400)
	earthDaysInYear := 365.25
	return ageInSeconds / (planetOrbitalPercentage * earthDaysInYear * secondsInDay)
}
