# TP-1 - GoLang

## Description
Membres du groupe : 
- Julien CHAZAL
- Louis LUBINEAU
- Florent PARIS

## Consignes
Créer un annuaire :
- au minimum “nom prénom” + “numéro de téléphone”
- Lister
- ajouter un contact
- supprimer
- modifier
- Vérifier si une personne avec le même nom existe déjà
- Test unitaires obligatoires

## Commandes
- `go run main.go --action add --name "Alice" --tel "0123456789"`
- `go run main.go --action delete --name "Alice"`
- `go run main.go --action edit --name "Alice" --tel "0987654321"`
- `go run main.go --action search --name "Alice"`
- `go run main.go --action list`
- `go run main.go --action search --name "Bob"`
