package main

import (
	"fmt"
)

type Personnage struct {
	nom        string
	vie        int
	puissance  int
	mort       bool
	inventaire [3]string
}

func main() {
	// initialisation de ma structure Personnage
	var p1 Personnage
	var p2 Personnage

	valeurParDefaut(&p1)
	valeurParDefaut(&p2)

	p1.nom = "barbare"
	p2.nom = "magicien"

	p1.affichage()
	p2.affichage()
}

// déclaration de ma méthode affichage() lié à ma structure Personnage
func (p Personnage) affichage() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Vie du personnage", p.nom, ":", p.vie)
	fmt.Println("Puissance du personnage", p.nom, ":", p.puissance)

	if p.mort {
		fmt.Println("Vie du personnage", p.nom, "est mort")
	} else {
		fmt.Println("Vie du personnage", p.nom, "est vivant")
	}

	fmt.Println("\nLe personnage", p.nom, "possède dans son inventaire :", p.vie)

	for _, item := range p.inventaire {
		fmt.Println("-", item)
	}
}

/*
@description : Initialise des valeurs par défaut pour un personnage

@return: void
*/
func valeurParDefaut(p1 *Personnage) {
	p1.nom = "inconnu"
	p1.vie = 50
	p1.puissance = 10
	p1.mort = false
	p1.inventaire = [3]string{"vide", "vide", "vide"}
}
