package main

import "testing"

func TestCheckGuess(t *testing.T) {
	if CheckGuess(5, 10) != "LOW" {
		t.Error("Deveria ser LOW")
	}
	if CheckGuess(15, 10) != "HIGH" {
		t.Error("Deveria ser HIGH")
	}
	if CheckGuess(10, 10) != "WIN" {
		t.Error("Deveria ser WIN")
	}
}
