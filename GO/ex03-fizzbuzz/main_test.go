package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	if GetFizzBuzz(3) != "Fizz" {
		t.Error("Erro no 3")
	}
	if GetFizzBuzz(5) != "Buzz" {
		t.Error("Erro no 5")
	}
	if GetFizzBuzz(15) != "FizzBuzz" {
		t.Error("Erro no 15")
	}
}
