package main

import "testing"

func TestCalculos(t *testing.T) {
	// Testando Celsius para Fahrenheit
	res1 := CelsiusToFahrenheit(0)
	if res1 != 32 {
		t.Errorf("Erro no Celsius! Esperava 32, recebi %.2f", res1)
	}

	// Testando Fahrenheit para Celsius
	res2 := FahrenheitToCelsius(212)
	if res2 != 100 {
		t.Errorf("Erro no Fahrenheit! Esperava 100, recebi %.2f", res2)
	}
}
