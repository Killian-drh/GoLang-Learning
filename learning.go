package main

import (
	"fmt"
)

func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

    // récupération de l'index et de la valeur
	for index, j := range jours {
		fmt.Println(j, "est le jour numéro", (index + 1), "de la semaine")
	}
}
