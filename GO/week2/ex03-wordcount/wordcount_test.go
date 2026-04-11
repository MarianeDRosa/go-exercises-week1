package main

import "testing"

func TestCountWordsSorted(t *testing.T) {
	text := "apple banana apple orange apple banana"
	results := CountWords(text)

	// In this input, "apple" appears 3 times, so it MUST be first
	if results[0].Word != "apple" {
		t.Errorf("Expected first word to be 'apple', got %s", results[0].Word)
	}

	if results[0].Count != 3 {
		t.Errorf("Expected 'apple' to have count 3, got %d", results[0].Count)
	}

	// "banana" appears 2 times, so it should be second
	if results[1].Word != "banana" {
		t.Errorf("Expected second word to be 'banana', got %s", results[1].Word)
	}
}
