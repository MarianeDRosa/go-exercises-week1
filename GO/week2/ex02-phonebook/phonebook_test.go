package main

import "testing"

func TestPhonebook(t *testing.T) {
	// Limpa a agenda antes de testar
	Phonebook = []Contact{}

	// 1. Teste Add
	AddContact(Contact{Name: "Davi", Email: "davi@test.com"})
	if len(Phonebook) != 1 {
		t.Errorf("Deveria ter 1 contato, tem %d", len(Phonebook))
	}

	// 2. Teste Search
	found := SearchByName("Davi")
	if found == nil || found.Name != "Davi" {
		t.Error("Não encontrou o contato Davi")
	}

	// 3. Teste Delete
	deleted := DeleteContact("Davi")
	if !deleted || len(Phonebook) != 0 {
		t.Error("Falhou ao deletar o contato")
	}
}
