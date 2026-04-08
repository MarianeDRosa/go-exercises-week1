package main

import "testing"

func TestEvenOdd(t *testing.T) {
	if IsEven(10) != true {
		t.Error("10 deveria ser par")
	}
	if IsEven(7) != false {
		t.Error("7 deveria ser ímpar")
	}
}
