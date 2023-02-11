package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	return (value - 32) * (5.0/9)
}

func CelsiusToFahrenheit(value float64) float64 {
	return (value*1.8) + 32
}

// De andre konverteringsfunksjonene implementere her
// ...