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

func TestSearchContact(t *testing.T) {
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

	contact, err := searchContact("Alice")
	if err != nil {
		t.Fatalf("Erreur lors de la recherche du contact : %v", err)
	}
	if contact == nil {
		t.Error("Le contact 'Alice' n'a pas été trouvé.")
	}
	if contact.Name != "Alice" {
		t.Errorf("Le nom du contact trouvé est incorrect : %s", contact.Name)
	}
}

func TestDeleteContact(t *testing.T) {
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

	err = deleteContact("Alice")
	if err != nil {
		t.Fatalf("Erreur lors de la suppression du contact : %v", err)
	}

	_, err = searchContact("Alice")
	if err == nil {
		t.Error("Le contact 'Alice' n'a pas été supprimé.")
	}
}
