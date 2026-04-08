package main

import "testing"

func TestContactLogic(t *testing.T) {
	// 1. Testa a criação
	c := CreateContact("Davi", "davi@test.com", "12345")
	if c.Name != "Davi" || c.Email != "davi@test.com" {
		t.Errorf("Falha ao criar contato!")
	}

	// 2. Testa a atualização do email (Ponteiro)
	c.UpdateEmail("davi.novo@test.com")
	if c.Email != "davi.novo@test.com" {
		t.Errorf("O email não foi atualizado corretamente. Recebi: %s", c.Email)
	}
}
