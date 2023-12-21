// Package weather provides an api to provide a forecast
// provided a city, and the condition in that city.
package weather

// CurrentCondition is the weather condition.
var CurrentCondition string

// CurrentLocation is the city name.
var CurrentLocation string

// Forecast provides a message for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
