package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculate recebe uma string "num1 op num2" e resolve
func Calculate(expression string) (float64, error) {
	// 1. Quebra a string em espaços
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, fmt.Errorf("formato inválido! Use: '3 + 5'")
	}

	// 2. Converte os números de string para float64
	n1, err1 := strconv.ParseFloat(parts[0], 64)
	n2, err2 := strconv.ParseFloat(parts[2], 64)
	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("números inválidos")
	}

	// 3. Verifica o operador
	operator := parts[1]
	switch operator {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, fmt.Errorf("divisão por zero não permitida")
		}
		return n1 / n2, nil
	default:
		return 0, fmt.Errorf("operador desconhecido: %s", operator)
	}
}

func main() {
	res, err := Calculate("10 / 2")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", res)
	}
}
