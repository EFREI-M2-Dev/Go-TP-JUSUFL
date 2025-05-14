package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Action to perform: add, list, update, delete, search")
	name := flag.String("name", "", "Name of the contact")
	newname := flag.String("newname", "", "New name of the contact")
	number := flag.String("number", "", "Phone number of the contact")
	flag.Parse()

	_, err := loadPhonebook()
	if err != nil {
		fmt.Println("Error during loading :", err)
		return
	}

	switch *action {
	case "add":
		fmt.Printf("-- Ajout d’un contact --\n")
		if *name == "" || *number == "" {
			fmt.Println("Veuillez fournir un nom et un numéro (--name et --number).")
			return
		}

		err := addContact(*name, *number)
		if err != nil {
			fmt.Println("Erreur lors de l’ajout :", err)
			return
		}
		fmt.Printf("Contact ajouté : %s (%s)\n", *name, *number)
	case "list":
		fmt.Printf("-- Liste des contacts -- \n")
		displayContacts()
	case "search":
		fmt.Printf("-- Recherche d'un contact --\n")
		if *name == "" {
			fmt.Println("Please provide a name to search.")
			return
		}

		contact, err := searchContact(*name)
		if err != nil {
			return
		}

		if contact != nil {
			fmt.Printf("Nom : %s\n", contact.Name)
			fmt.Printf("Téléphone : %s\n", contact.Number)
		}
	case "update":
		if *name == "" || *number == "" || *newname == "" {
			fmt.Println("Please provide a name, newname and a number to update.")
			return
		}
		err := updateContact(*name, *newname, *number)
		if err != nil {
			fmt.Println("Error during update :", err)
			return
		} else {
			fmt.Printf("Contact %s updated to %s\n", *name, *number)
		}
	case "delete":
		fmt.Printf("-- Suppression du contact --\n")
		err := deleteContact(*name)
		if err != nil {
			fmt.Println("Error during deletion :", err)
			return
		}
	default:
		fmt.Println("Action inconnue")
	}
}
