package main

import "testing"

func TestGetGreeting(t *testing.T) {
	expected := "Hello, World!"
	result := GetGreeting()

	if result != expected {
		t.Errorf("Erro no Hello World! Esperava '%s', mas recebi '%s'", expected, result)
	}
}
