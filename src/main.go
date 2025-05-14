package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "Action to perform: add, list, update, delete")
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
		fmt.Printf("Liste des contacts :\n")
		displayContacts()
	case "update":
		fmt.Printf("update contact")
	case "delete":
		fmt.Printf("delete contact")
	default:
		fmt.Println("Action inconnue")
	}
}
