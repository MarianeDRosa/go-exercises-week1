package main

import "fmt"

// GetGreeting devolve a frase para o teste conferir
func GetGreeting() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(GetGreeting())
}
