package main

import "testing"

func TestCalculate(t *testing.T) {
	// Teste de Soma
	res, _ := Calculate("5 + 5")
	if res != 10 {
		t.Errorf("Soma errada! Esperava 10, recebi %.2f", res)
	}

	// Teste de Divisão
	res, _ = Calculate("10 / 2")
	if res != 5 {
		t.Errorf("Divisão errada! Esperava 5, recebi %.2f", res)
	}

	// TESTE DE QA: Divisão por zero (Deve retornar erro)
	_, err := Calculate("10 / 0")
	if err == nil {
		t.Error("Deveria ter retornado erro ao dividir por zero!")
	}
}
