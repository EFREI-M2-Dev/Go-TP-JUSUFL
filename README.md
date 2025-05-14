# TP-1 - GoLang

## Description

Membres du groupe :

- Julien CHAZAL
- Louis LUBINEAU
- Florent PARIS

## Consignes

Créer un annuaire :

- Au minimum “nom prénom” + “numéro de téléphone”
- Lister
- Ajouter un contact
- Supprimer
- Modifier
- Vérifier si une personne avec le même nom existe déjà
- Test unitaires obligatoires

## Commandes

- `go run . --action list`
- `go run . --action add --name "Francis Huster" --number "0123456789"`
- `go run . --action update --name "Francis Huster" --newName "Marc Lavoine" --number "118218"`
- `go run . --action search --name "Francis Huster"`
- `go run . --action delete --name "Francis Huster"`
