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
