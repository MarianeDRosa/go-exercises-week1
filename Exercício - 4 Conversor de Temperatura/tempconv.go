package main

import "fmt"

func main() {
	var escolha int
	var temp float64

	fmt.Println("Escolha a conversão:")
	fmt.Println("1 - Celsius para Fahrenheit")
	fmt.Println("2 - Fahrenheit para Celsius")
	fmt.Scan(&escolha)

	fmt.Print("Digite a temperatura: ")
	fmt.Scan(&temp)

	if escolha == 1 {
		res := (temp * 9 / 5) + 32
		fmt.Printf("%.2f°C é igual a %.2f°F\n", temp, res)
	} else if escolha == 2 {
		res := (temp - 32) * 5 / 9
		fmt.Printf("%.2f°F é igual a %.2f°C\n", temp, res)
	} else {
		fmt.Println("Opção inválida!")
	}
}
