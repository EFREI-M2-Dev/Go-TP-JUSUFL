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

func TestAddContact(t *testing.T) {
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

	err = addContact("Charlie", "91011")
	if err != nil {
		t.Fatalf("Erreur lors de l'ajout du contact : %v", err)
	}

	phonebook, err := loadPhonebook()
	if err != nil {
		t.Fatalf("Erreur lors du chargement du phonebook après ajout : %v", err)
	}

	var contactFound bool
	for _, c := range phonebook.Contacts {
		if c.Name == "Charlie" && c.Number == "91011" {
			contactFound = true
		}
	}

	if !contactFound {
		t.Error("Le contact 'Charlie' n'a pas été ajouté correctement.")
	}

	contactFound = false
	for _, c := range phonebook.Contacts {
		if c.Name == "Alice" && c.Number == "1234" {
			contactFound = true
		}
	}

	if !contactFound {
		t.Errorf("Le contact 'Alice' a été supprimé.")
	}

	contactFound = false
	for _, c := range phonebook.Contacts {
		if c.Name == "Bob" && c.Number == "5678" {
			contactFound = true
		}
	}

	if !contactFound {
		t.Errorf("Le contact 'Bob' a été supprimé.")
	}
}
