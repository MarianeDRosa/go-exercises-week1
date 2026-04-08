package main

import "fmt"

// FUNÇÕES QUE O TESTE VAI LER:
// ---------------------------------------------------------
func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// ---------------------------------------------------------

func main() {
	var choice int
	var temp float64

	fmt.Println("1: Celsius to Fahrenheit\n2: Fahrenheit to Celsius")
	fmt.Print("Choose: ")
	fmt.Scan(&choice)
	fmt.Print("Temperature: ")
	fmt.Scan(&temp)

	switch choice {
	case 1:
		// Aqui chamamos a função em vez de fazer a conta direto
		result := CelsiusToFahrenheit(temp)
		fmt.Printf("%.2f°C is %.2f°F\n", temp, result)
	case 2:
		result := FahrenheitToCelsius(temp)
		fmt.Printf("%.2f°F is %.2f°C\n", temp, result)
	default:
		fmt.Println("Invalid option!")
	}
}
