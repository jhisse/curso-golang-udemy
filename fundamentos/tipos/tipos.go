package main

import (
	"math"
	"reflect"
	"fmt"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200000000))

	// sem sinal (positivos). uint8, uint16, uint32, uint64
	var b byte = 255 // byte é um "apelido para uint8"
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal
	i1 := math.MaxUint32
	fmt.Println("O valor maximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	// rune
	var i2 rune = 'a' // unicode, é definido pelas aspas simples
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //False

	// String
	s1 := "Olá meu nome e José" // Os acentos irão refletir no tamanho da string
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// Strings com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é 
	José`
	fmt.Println("O tamanho da string é", len(s2))

	// char existe em go?
	char := 'a' // defini uma rune
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}