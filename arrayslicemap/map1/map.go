package main

import (
	"fmt"
)

func main() {
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"

	fmt.Println(aprovados)

	for id, nome := range aprovados {
		fmt.Printf("%s (ID: %d)\n", nome, id)
	}

	fmt.Println(aprovados[2])
	delete(aprovados, 1)
	fmt.Println(aprovados[1])
}
