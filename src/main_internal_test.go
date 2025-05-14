package main

import (
	"strings"
	"testing"
)

func TestDisplayContact(t *testing.T) {
	testData := PhoneBook{
		Contacts: []Contact{
			{Name: "Alice", Number: "1234"},
			{Name: "Bob", Number: "5678"},
		},
	}

	err := savePhonebook(testData)
	if err != nil {
		t.Fatalf("Erreur lors de la sauvegarde du phonebook : %v", err)
	}

	output := displayContacts()

	if output == "" {
		t.Error("Aucun contact n'a été affiché.")
	}
	if !contains(output, "Alice") {
		t.Errorf("Le contact 'Alice' est manquant dans la sortie.")
	}
	if !contains(output, "Bob") {
		t.Errorf("Le contact 'Bob' est manquant dans la sortie.")
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
