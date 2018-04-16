package main

import (
	"fmt"
)

func obterNota(nota float64) string {
	if(nota >= 6) {
		return "Aprovado"
	}
	return "Reprovado"
	// return "aprovado" ? nota >= 6 : "reporvado" // Não existe operador ternário igual ao C
}

func main() {
	fmt.Println(obterNota(6.2))
}