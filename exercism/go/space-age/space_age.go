package space

type Planet string

const EarthYear = 31_557_600

var planetYears = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	year, ok := planetYears[string(planet)]
	if !ok {
		return -1
	}
	return seconds / (EarthYear * year)
}
