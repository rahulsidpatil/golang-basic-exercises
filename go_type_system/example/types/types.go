package main

import (
	"fmt"
)

func main() {
	type (
		// Temperature => defined/named type; float64 => Source type; float64 => Underlying type
		Temperature float64

		// Celsius=> defined/named type; Temperature => Source type; float64 => Underlying type
		Celsius Temperature

		// Fahrenheit => defined/named type; Temperature => Source type; float64 => Underlying type
		Fahrenheit Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)
}
