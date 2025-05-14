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

func displayContacts() {
	phoneBook, err := loadPhonebook()
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	if len(phoneBook.Contacts) == 0 {
		fmt.Println("Phonebook is empty.")
		return
	}
	for _, c := range phoneBook.Contacts {
		fmt.Printf("Nom: %s, Téléphone: %s\n", c.Name, c.Number)
	}
}

func savePhonebook(phonebook PhoneBook) error {
	data, err := json.MarshalIndent(phonebook, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
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

func deleteContact(name string) error {
	phonebook, err := loadPhonebook()
	if err != nil {
		return err
	}

	found := false
	for i, c := range phonebook.Contacts {
		if c.Name == name {
			phonebook.Contacts = append(phonebook.Contacts[:i], phonebook.Contacts[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("le contact '%s' n'existe pas", name)
	}

	return savePhonebook(phonebook)
}