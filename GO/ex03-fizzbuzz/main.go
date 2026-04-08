package main

import "fmt"

// Lógica isolada para o teste
func GetFizzBuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", i)
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(GetFizzBuzz(i))
	}
}
