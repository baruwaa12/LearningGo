package spaceage

import (
	"fmt"
	"os"
)

// Planet represents a planet within our solar system. NOTE: This does not include dwarf planets, planetoids or other celestial bodies but
type Planet string


// Given a number of seconds and a specific planet calculates a persons age in earth years.
func Age(seconds float64, planet Planet) float64 {

	var years float64


	switch planet {
	case "Mercury":
		years = seconds / 7600543.81992 // 0.240846 earth years converted to seconds. See https://en.wikipedia.org/wiki/Mercury_(planet)
	case "Venus":
		years = seconds / 19414172.4048 // 0.61519726 earth years converted to seconds. See https://en.wikipedia.org/wiki/Venus
	case "Earth":
		years = seconds / 31557600 //
	case "Mars":
		years = seconds / 59354032.6901 // 1.8808158 earth years converted to seconds. See https://en.wikipedia.org/wiki/Mars
	case "Jupiter":
		years = seconds / 374355659.124 // 11.862615 earth years converted to seconds. See https://en.wikipedia.org/wiki/Jupiter
	case "Saturn":
		years = seconds / 929292362.885 // 29.447498 earth years converted to seconds. See https://en.wikipedia.org/wiki/Saturn
	case "Uranus":
		years = seconds / 2651370019.33 // 84.016846 earth years converted to seconds. See https://en.wikipedia.org/wiki/Uranus
	case "Neptune":
		years = seconds / 5200418560.03 // 164.79132 earth years converted to seconds. See https://en.wikipedia.org/wiki/Neptune
	default :
		fmt.Printf("Unexpected input %s\n", planet)
		os.Exit(1)
	}

	// Return the age
	return years
}