package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Nom string `json:"nom"`
	Tel string `json:"tel"`
}

type Annuaire struct {
	Contacts []Contact `json:"contacts"`
}

const storageFile = "annuaire.json"

func chargerAnnuaire() (Annuaire, error) {
	var a Annuaire
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

func listerContacts() {
	annuaire, err := chargerAnnuaire()
	if err != nil {
		fmt.Println("Erreur de chargement :", err)
		return
	}
	if len(annuaire.Contacts) == 0 {
		fmt.Println("Annuaire vide")
		return
	}
	for _, c := range annuaire.Contacts {
		fmt.Printf("Nom: %s, Téléphone: %s\n", c.Nom, c.Tel)
	}
}
