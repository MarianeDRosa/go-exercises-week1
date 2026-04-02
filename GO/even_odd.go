package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1]
	num, _ := strconv.Atoi(entrada)

	if num%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}

	if num > 0 {
		fmt.Println("Positivo")
	} else if num < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Zero")
	}

	absoluto := num
	if num < 0 {
		absoluto = num * -1
	}
	fmt.Printf("Absoluto: %d\n", absoluto)
}
