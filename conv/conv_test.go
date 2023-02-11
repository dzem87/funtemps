package conv

import (
	"reflect"
	"testing"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := FarhenheitToCelsius(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

	func TestCelsiusToFahrenheit(t *testing.T) {
		type test struct { //lager en klasse
			input float64 //variabel for input med verdi	
			want  float64
		}
	
		testCtoF := []test{
			{input: 30, want: 86}, //gir verdi vil ha verdi
		}
	
		for _, cf := range testCtoF { //for loop som går gjennom alle verdiene til tests
			got := CelsiusToFahrenheit(cf.input) //got = input
			if !reflect.DeepEqual(cf.want, got) { //sjekker om want er lik got
				t.Errorf("expected: %v, got: %v", cf.want, got)
			}
		}
	}



// De andre testfunksjonene implementeres her
// ...