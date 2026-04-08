package main

import (
	"fmt"
	"math/rand/v2"
)

func CheckGuess(guess, target int) string {
	if guess < target {
		return "LOW"
	}
	if guess > target {
		return "HIGH"
	}
	return "WIN"
}

func main() {
	target := rand.IntN(100) + 1
	var guess int
	fmt.Println("Adivinhe o número!")

	for {
		fmt.Scan(&guess)
		result := CheckGuess(guess, target)
		if result == "WIN" {
			fmt.Println("Acertou!")
			break
		}
		fmt.Println(result)
	}
}
