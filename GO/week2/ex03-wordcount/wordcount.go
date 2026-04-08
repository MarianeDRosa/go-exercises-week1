package main

import (
	"fmt"
	"strings"
)

// CountWords recebe uma string e devolve um Map com as contagens
func CountWords(text string) map[string]int {
	// 1. Criamos um mapa vazio
	counts := make(map[string]int)

	// 2. Quebramos o texto em palavras (separadas por espaço)
	words := strings.Fields(strings.ToLower(text))

	// 3. Loop para contar
	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	frase := "Go is amazing and Go is fast"
	resultado := CountWords(frase)
	fmt.Println(resultado)
}
