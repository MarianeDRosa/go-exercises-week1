package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <número>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Por favor, digite um número válido.")
		return
	}

	if IsEven(num) {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}
}
