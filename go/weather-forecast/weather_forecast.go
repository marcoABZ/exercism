// Package weather provides a Forecast function to get a
// current wheather condition on a given location.
package weather

// CurrentCondition is a string that represents the current wheather condition.
var CurrentCondition string

// CurrentLocation is a string that represents the current location where the weather is being inspected.
var CurrentLocation string

// Forecast receives the name of a city and the current wheather condition and returns an string with the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
