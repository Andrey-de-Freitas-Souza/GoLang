package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "Futebol"
	switch esporteFavorito{
		
	case "Futebol":
		fmt.Println("O esporte favorito é Futebol")
	case "Vôlei":
		fmt.Println("O esporte favorito é Vôlei")
	case "Basquete":
		fmt.Println("O esporte favorito é Basquete")
	case "Natação":
		fmt.Println("O esporte favorito é Natação")

	}
}
