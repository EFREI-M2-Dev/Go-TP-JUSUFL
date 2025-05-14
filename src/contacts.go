package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name   string `json:"nom"`
	Number string `json:"tel"`
}

type PhoneBook struct {
	Contacts []Contact `json:"contacts"`
}

const storageFile = "annuaire.json"

func loadPhonebook() (PhoneBook, error) {
	var a PhoneBook
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return a, nil
		}
		return a, err
	}
	err = json.Unmarshal(data, &a)
	return a, err
}

func displayContacts() string {
	phoneBook, err := loadPhonebook()
	if err != nil {
		return fmt.Sprintf("Error: %v\n", err)
	}
	if len(phoneBook.Contacts) == 0 {
		return "Contacts list is empty."
	}
	var result string
	for _, c := range phoneBook.Contacts {
		result += fmt.Sprintf("Nom: %s, Téléphone: %s\n", c.Name, c.Number)
	}
	return result
}

func savePhonebook(phonebook PhoneBook) error {
	data, err := json.MarshalIndent(phonebook, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func addContact(name, number string) error {
	phonebook, err := loadPhonebook()
	if err != nil {
		return fmt.Errorf("error while loading phonebook: %w", err)
	}

	for _, c := range phonebook.Contacts {
		if c.Name == name {
			return fmt.Errorf("le contact '%s' existe déjà", name)
		}
	}

	newContact := Contact{
		Name:   name,
		Number: number,
	}

	phonebook.Contacts = append(phonebook.Contacts, newContact)

	err = savePhonebook(phonebook)
	if err != nil {
		return fmt.Errorf("error while saving phonebook: %w", err)
	}

	return nil
}

func updateContact(name, newName, newNumber string) error {
	phonebook, err := loadPhonebook()
	if err != nil {
		return err
	}

	for i, c := range phonebook.Contacts {
		if c.Name == name {
			phonebook.Contacts[i].Name = newName
			phonebook.Contacts[i].Number = newNumber
			return savePhonebook(phonebook)
		}
	}

	return fmt.Errorf("contact not found")
}

func searchContact(name string) (*Contact, error) {
	phonebook, err := loadPhonebook()
	if err != nil {
		return nil, err
	}

	for _, c := range phonebook.Contacts {
		if c.Name == name {
			return &c, err
		}
	}

	return nil, fmt.Errorf("le contact '%s' n'existe pas", name)
}

func deleteContact(name string) error {
	_, err := searchContact(name)
	if err != nil {
		return err
	}

	phonebook, err := loadPhonebook()
	if err != nil {
		return err
	}

	newContacts := make([]Contact, 0, len(phonebook.Contacts))
	for _, c := range phonebook.Contacts {
		if c.Name != name {
			newContacts = append(newContacts, c)
		}
	}

	phonebook.Contacts = newContacts
	return savePhonebook(phonebook)
}
