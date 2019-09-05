package space

import "fmt"

// Planet stores a type of planet
type Planet string

// Age takes an age in seconds and a planet.  It returns an age in years
// on the given planet.
func Age(ageInSeconds float64, planet Planet) float64 {
	//The following variables are percentage of earth year.
	mercuryOrbitalPercentage := 0.2408467
	venusOrbitalPercentage := 0.61519726
	marsOrbitalPercentage := 1.8808158
	jupiterOrbitalPercentage := 11.862615
	saturnOrbitalPercentage := 29.447498
	uranusOrbitalPercentage := 84.016846
	neptuneOrbitalPercentage := 164.79132

	switch planet {
	case "Earth":
		return determineRelativeAge(ageInSeconds, 1)
	case "Mercury":
		return determineRelativeAge(ageInSeconds, mercuryOrbitalPercentage)
	case "Venus":
		return determineRelativeAge(ageInSeconds, venusOrbitalPercentage)
	case "Mars":
		return determineRelativeAge(ageInSeconds, marsOrbitalPercentage)
	case "Jupiter":
		return determineRelativeAge(ageInSeconds, jupiterOrbitalPercentage)
	case "Saturn":
		return determineRelativeAge(ageInSeconds, saturnOrbitalPercentage)
	case "Uranus":
		return determineRelativeAge(ageInSeconds, uranusOrbitalPercentage)
	case "Neptune":
		return determineRelativeAge(ageInSeconds, neptuneOrbitalPercentage)
	default:
		panic(fmt.Sprintf("Unrecognized planet: %v", planet))
	}
}

func determineRelativeAge(ageInSeconds float64, planetOrbitalPercentage float64) float64 {
	return ageInSeconds / calculateSecondsInYear(planetOrbitalPercentage)
}

func calculateSecondsInYear(earthOrbitalPercent float64) float64 {
	secondsInDay := float64(86400)
	earthDaysInYear := 365.25

	return earthOrbitalPercent * earthDaysInYear * secondsInDay
}
