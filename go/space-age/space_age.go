package space

import "math"

type Planet string

const (
	secondMinuteHourByDay  = 86400.0
	daysByEarthYear        = 365.25
	mercuryDaysByEarthYear = daysByEarthYear * 0.2408467
	venusDaysByEarthYear   = daysByEarthYear * 0.61519726
	marsDaysByEarthYear    = daysByEarthYear * 1.8808158
	jupiterDaysByEarthYear = daysByEarthYear * 11.862615
	saturnDaysByEarthYear  = daysByEarthYear * 29.447498
	uranusDaysByEarthYear  = daysByEarthYear * 84.016846
	neptuneDaysByEarhYear  = daysByEarthYear * 164.79132
)

func Age(seconds float64, planet Planet) float64 {
	var age float64
	switch planet {
	case "Mercury":
		age = secondMinuteHourByDay * mercuryDaysByEarthYear
	case "Venus":
		age = secondMinuteHourByDay * venusDaysByEarthYear
	case "Earth":
		age = secondMinuteHourByDay * daysByEarthYear
	case "Mars":
		age = secondMinuteHourByDay * marsDaysByEarthYear
	case "Jupiter":
		age = secondMinuteHourByDay * jupiterDaysByEarthYear
	case "Saturn":
		age = secondMinuteHourByDay * saturnDaysByEarthYear
	case "Uranus":
		age = secondMinuteHourByDay * uranusDaysByEarthYear
	case "Neptune":
		age = secondMinuteHourByDay * neptuneDaysByEarhYear
	default:
		return math.NaN()
	}
	return seconds / age
}
