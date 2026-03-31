package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	alvo := rand.Intn(100) + 1
	var chute int
	fmt.Println(" === JOGO DE ADIVINHAÇÃO === ")
	fmt.Println("Tente adivinhar o número entre 1 e 100.")

	for {
		fmt.Print("Seu chute: ")
		fmt.Scan(&chute)
		if chute < alvo {
			fmt.Println("Muito BAIXO!")
		} else if chute > alvo {
			fmt.Println("Muito ALTO!")
		} else {
			fmt.Println("ACERTOU!")
			break

		}
	}
}
