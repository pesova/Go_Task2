package main

import (
	"fmt"
	"math"
)

//Converter contains all the units as fields
type Converter struct {
	Minutes
	Seconds
	Centimeters
	Feet
	Celsius
	Fahrenheit
	Radian
	Degrees
	Kilograms
	Pounds
}
type Fahrenheit float64

type Radian float64

type Degrees float64

type Kilograms float64

type Pounds float64

type Minutes float64

type Seconds float64

type Centimeters float64

type Feet float64

type Celsius float64

//MinutesToSeconds ...
func (c Converter) MinutesToSeconds(v Minutes) Seconds {
	c.Seconds = Seconds(v) * 60
	return c.Seconds
}

//SecondsToMinutes ...
func (c Converter) SecondsToMinutes(v Seconds) Minutes {
	return Minutes(v / 60)
}

//FeetToCentimeters ...
func (c Converter) FeetToCentimeters(v Feet) Centimeters {
	return Centimeters(v * 30.48)
}

//CentimetersToFeet ...
func (c Converter) CentimetersToFeet(v Centimeters) Feet {
	return Feet(v / 30.48)
}

//FahrenheitToCelsius ...
func (c Converter) FahrenheitToCelsius(v Fahrenheit) Celsius {
	return Celsius((v - 32) * 5 / 9)
}

//CelsiusToFahrenheit ...
func (c Converter) CelsiusToFahrenheit(v Celsius) Fahrenheit {
	return Fahrenheit(v*9/5 + 32)
}

//RadianToDegrees ...
func (c Converter) RadianToDegrees(v Radian) Degrees {
	return Degrees(v * (180 / math.Pi))
}

//DegreesToRadian ...
func (c Converter) DegreesToRadian(v Degrees) Radian {
	return Radian(v * (math.Pi / 180))
}

//KilogramsToPounds ...
func (c Converter) KilogramsToPounds(v Kilograms) Pounds {
	return Pounds(v * 2.205)
}

//PoundsToKilograms ...
func (c Converter) PoundsToKilograms(v Pounds) Kilograms {
	return Kilograms(v / 2.205)
}

func main() {
	c := new(Converter)
	fmt.Printf("1 minute is equal to %v Seconds\n", c.MinutesToSeconds(1.0))
	fmt.Printf("1 second is equal to %v Minutes\n", c.SecondsToMinutes(1.0))
	fmt.Printf("1 centimeter is equal to %v Feet\n", c.CentimetersToFeet(1.0))
	fmt.Printf("1 foot is equal to %v Centimeters\n", c.FeetToCentimeters(1.0))
	fmt.Printf("1 celsius is equal to %v Fahrenheit\n", c.CelsiusToFahrenheit(1.0))
	fmt.Printf("1 fahrenheit is equal to %v Celsuis\n", c.FahrenheitToCelsius(1.0))
	fmt.Printf("1 radian is equal to %v Degrees\n", c.RadianToDegrees(1.0))
	fmt.Printf("1 degree is equal to %v radian\n", c.DegreesToRadian(1.0))
	fmt.Printf("1 kilogram is equal to %v pounds\n", c.KilogramsToPounds(1.0))
	fmt.Printf("1 pound is equal to %v kilograms\n", c.PoundsToKilograms(1.0))
}
