package main

import (
	"strconv"
	"fmt"
)

func main() {
	x:= 2.4
	y :=  2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste", string(97)) // Não converte int em string e sim o correspondente da tabela unicode

	fmt.Println("Teste", strconv.Itoa(97))

	num, _ := strconv.Atoi("97")
	fmt.Println("Teste", num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}