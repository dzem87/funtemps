package conv

import "fmt"

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) string {
	celsius := (value - 32) * (5.0/9)
  return fmt.Sprintf("%.2f", celsius)

}

func FahrenheitToKelvin(value float64) string {
  kelvin := ((value - 32) * (5.0/9)) + 273.15
  return fmt.Sprintf("%.2f", kelvin)
}

func CelsiusToFahrenheit(value float64) string {
	fahrenheit := (value*1.8) + 32
  return fmt.Sprintf("%.2f", fahrenheit)
}

func CelsiusToKelvin(value float64) string {
	fahrenheit := value + 273.15
  return fmt.Sprintf("%.2f", fahrenheit)
}

func KelvinToCelsius(value float64) string {
	celsius := value - 273.15
  return fmt.Sprintf("%.2f", celsius)
}

func KelvinToFahrenheit(value float64) string {
	fahrenheit := ((value - 273.15)*(9.0/5)) + 32
  return fmt.Sprintf("%.2f", fahrenheit)
}

