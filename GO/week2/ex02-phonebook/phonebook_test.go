package main

import "testing"

func TestPhonebook(t *testing.T) {
	// CRITICAL: Reset the global slice to ensure a clean state for this test
	Phonebook = []Contact{}

	// 1. Test AddContact
	newContact := Contact{Name: "Mariane", Email: "mari@test.com"}
	AddContact(newContact)

	if len(Phonebook) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(Phonebook))
	}

	// 2. Test SearchByName
	found, err := SearchByName("Mariane")
	if err != nil || found.Name != "Mariane" {
		t.Errorf("Search failed: expected Mariane, got error or different name")
	}

	// 3. Test DeleteContact
	errDelete := DeleteContact("Mariane")
	if errDelete != nil || len(Phonebook) != 0 {
		t.Errorf("Delete failed: contact should have been removed")
	}
}
