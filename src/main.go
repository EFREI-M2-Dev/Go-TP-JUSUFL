package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Action to perform: add, list, update, delete")
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
		fmt.Printf("add contact")
	case "list":
		fmt.Printf("-- Liste des contacts -- \n")
		displayContacts()
	case "update":
		if *name == "" || *number == "" || *newname == "" {
			fmt.Println("Please provide a name, newname and a number to update.")
			return
		}
		err := updateContact(*name, *newname, *number)
		if err != nil {
			fmt.Println("Error during update :", err)
			return
		}else {
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
