package main

import "testing"

func TestCountWords(t *testing.T) {
	text := "teste teste QA"
	counts := CountWords(text)

	// Valida se "teste" apareceu 2 vezes
	if counts["teste"] != 2 {
		t.Errorf("Esperava 2 ocorrências de 'teste', recebi %d", counts["teste"])
	}

	// Valida se "qa" apareceu 1 vez
	if counts["qa"] != 1 {
		t.Errorf("Esperava 1 ocorrência de 'qa', recebi %d", counts["qa"])
	}
}
